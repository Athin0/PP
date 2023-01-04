package astParser

import (
	"PP/worker/asyncDispatching"
	"PP/worker/genericMath"
	"PP/worker/grammar/lexer"
	"PP/worker/grammar/parser"
	"PP/worker/sequenceRepo/PgRepo"
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

var ErrParsingString = errors.New("error parsing string")

func TestTraverse(t *testing.T) {
	text := " [a] "
	lex := lexer.New([]rune(text))

	if _bsr, errs := parser.Parse(lex); len(errs) != 0 {
		log.Printf("Err in Parser: %v", ErrParsingString)

		t.Error(ErrParsingString)
	} else {
		repo, err := PgRepo.InitDB()
		if err != nil {
			t.Errorf("error in init repo : %v", err)
			return
		}
		root, err := BuildAST(_bsr.GetRoot(), nil, repo)
		if err != nil {
			t.Error(err)
			return
		}
		disp := asyncDispatching.NewDispatcher(root)

		fmt.Println("ans:", asyncDispatching.Traverse(disp))
	}
}

func InitTestingSequence(path string) (*genericMath.FloatSequence, error) {
	seq := genericMath.FloatSequence{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}
		data := strings.Split(line, " ")
		for _, snum := range data {
			f, _ := strconv.ParseFloat(snum, 64)
			seq.Append(f)
		}
		if err != nil {
			break
		}
	}

	return &seq, nil
}
