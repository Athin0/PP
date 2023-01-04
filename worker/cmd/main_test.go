package main

import (
	"PP/worker/sequenceRepo/PgRepo"
	"fmt"
	"log"
	"testing"
)

func TestCalculateSequence(t *testing.T) {
	db, err := PgRepo.InitDB()
	if err != nil {
		log.Println("error of init db: ", err)
		return
	}
	sequence, err := CalculateSequence("[a]", db)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sequence)
}
