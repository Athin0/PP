package astParser

import (
	"PP/worker/asyncDispatching"
	"PP/worker/grammar/parser/bsr"
	"PP/worker/sequenceRepo"
	"fmt"
	"regexp"
)

/*
Constants with regex values to identify operation
*/
const (
	regexMean = `^{ [A-z]+ }$`
	regexVar  = `^\[ [A-z]+ \]$`
	regexStd  = `^< [A-z]+ >$`
)

/*
Get terminal symbol from node
*/
func getUnaryOperation(token string) asyncDispatching.UnaryFloatFunction {
	match, _ := regexp.MatchString(regexMean, token)
	if match {
		return getMean
	}
	match, _ = regexp.MatchString(regexVar, token)
	if match {
		return getVariance
	}
	match, _ = regexp.MatchString(regexStd, token)
	if match {
		return getStandardDeviation
	}
	return nil
}
func getBinaryOperation(token string) (binAct asyncDispatching.BinaryFloatFunction) {
	if token == "+" {
		binAct = add
	} else if token == "-" {
		binAct = sub
	} else if token == "*" {
		binAct = mul
	} else if token == "/" {
		binAct = div
	} else if token == "!" {
		binAct = min
	} else if token == "?" {
		binAct = max
	}
	return binAct
}

/*
Function to build AST that is usable by Dispatcher
Returns root of AST
*/
func BuildAST(bsr bsr.BSR, prev *asyncDispatching.Node, dataRepo sequenceRepo.IRepo) (*asyncDispatching.Node, error) {
	if len(bsr.GetAllNTChildren()) == 3 {
		var binAct asyncDispatching.BinaryFloatFunction
		token := bsr.GetNTChildrenI(1)[0].Label.Symbols().String()
		binAct = getBinaryOperation(token)
		if binAct == nil {
			return nil, fmt.Errorf("no such bin operation: %s", token)
		}
		root := &asyncDispatching.Node{
			BinaryAction: binAct,
		}
		var err error
		root.Left, err = BuildAST(bsr.GetNTChildrenI(0)[0], root, dataRepo)
		root.Right, err = BuildAST(bsr.GetNTChildrenI(2)[0], root, dataRepo)
		if err != nil {
			return nil, err
		}
		return root, nil
	} else if len(bsr.GetAllNTChildren()) == 1 {

		token := bsr.GetAllNTChildren()[0][0].Label.Symbols().String()
		operation := getUnaryOperation(token)
		if operation == nil {
			return BuildAST(bsr.GetNTChildrenI(0)[0], prev, dataRepo)
		}
		sym := bsr.GetNTChildI(0).GetTChildI(1).LiteralString()
		seq, err := dataRepo.GetSequence(sym)
		if err != nil {
			fmt.Printf("error in get sequence: %v", err)
			return nil, err
		}
		return &asyncDispatching.Node{
			InitialSequence: seq,
			Parent:          prev,
			UnaryAction:     operation,
		}, nil
	}

	return nil, nil
}
