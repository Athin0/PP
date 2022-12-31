package mthprsr

import (
	"log"

	"PP/app/internal/grammar/lexer"
	"PP/app/internal/grammar/parser"
)

func ValidateString(in string) bool {
	lex := lexer.New([]rune(in))

	if _, errs := parser.Parse(lex); len(errs) != 0 {
		log.Println("Error validating string")

		return false
	} else {
		return true
	}
}
