package astParser

import (
	"regexp"
	"unicode"

	"PP/worker/asyncDispatching"
	"PP/worker/genericMath"
	"PP/worker/grammar/parser/bsr"
)

/*
*
Constants with regex values to identify operation
*/
const (
	regexMean = `^{[A-z]+}$`
	regexVar  = `^\[[A-z]+\]$`
	regexStd  = `^<[A-z]+>$`
)

/*
*
Function substitutes
*/
var getMean = func(sequence *genericMath.FloatSequence) float64 {
	res, _ := sequence.GetMean()
	return res
}

var getVariance = func(sequence *genericMath.FloatSequence) float64 {
	res, _ := sequence.GetVariance()
	return res
}

var getStandardDeviation = func(sequence *genericMath.FloatSequence) float64 {
	res, _ := sequence.GetStandardDeviation()
	return res
}

var add = func(a, b float64) float64 {
	return genericMath.Sum(a, b)
}

var sub = func(a, b float64) float64 {
	return genericMath.Sum(a, -b)
}

var mul = func(a, b float64) float64 {
	return genericMath.Multiply(a, b)
}

var div = func(a, b float64) float64 {
	res, _ := genericMath.Divide(a, b)
	return res
}

var min = func(a, b float64) float64 {
	return genericMath.Min(a, b)
}

var max = func(a, b float64) float64 {
	return genericMath.Max(a, b)
}

/*
*
Get terminal symbol from node
*/
func parseBSRString(s []rune) string {
	for i := range s {
		if unicode.IsDigit(s[i]) && s[i+1] == ' ' && s[i+2] == '-' && s[i+3] == ' ' { //todo
			return string(s[i+4:])
		}
	}
	return ""
}

/*
Function to build AST that is usable by Dispatcher
Returns root of AST
*/
func BuildAST(bsr bsr.BSR, prev *asyncDispatching.Node) *asyncDispatching.Node {
	if len(bsr.GetAllNTChildren()) == 3 {
		var binAct asyncDispatching.BinaryFloatFunction
		token := parseBSRString([]rune(bsr.GetNTChildrenI(1)[0].String()))

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

		root := &asyncDispatching.Node{
			BinaryAction: binAct,
		}
		root.Left = BuildAST(bsr.GetNTChildrenI(0)[0], root)
		root.Right = BuildAST(bsr.GetNTChildrenI(2)[0], root)

		return root
	} else if len(bsr.GetAllNTChildren()) == 1 {
		token := parseBSRString([]rune(bsr.GetAllNTChildren()[0][0].String())) //todo может сразу в стринге передавать?

		match, _ := regexp.MatchString(regexMean, token) //todo обработать ошибки?
		if match {
			return &asyncDispatching.Node{
				Parent:      prev,
				UnaryAction: getMean,
			}
		}
		match, _ = regexp.MatchString(regexVar, token)
		if match {
			return &asyncDispatching.Node{
				Parent:      prev,
				UnaryAction: getVariance,
			}
		}
		match, _ = regexp.MatchString(regexStd, token)
		if match {
			return &asyncDispatching.Node{
				Parent:      prev,
				UnaryAction: getStandardDeviation,
			}
		}

		return BuildAST(bsr.GetNTChildrenI(0)[0], prev)
	}

	return nil
}
