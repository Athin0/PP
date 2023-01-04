package PgRepo

import (
	"PP/worker/sequenceRepo"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestPG(t *testing.T) {
	pg, _ := InitDB()
	seq, err := pg.GetSequence("a")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(seq)
}

func TestJSON(t *testing.T) {
	file, err := os.Open("../../static/data4.json")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)

	seq, err := sequenceRepo.GetSequenceJson(byteValue)

	if err != nil {
		return
	}
	println(seq.GetLength())
}

func TestAddJSON(t *testing.T) {
	pg, _ := InitDB()
	for _, data := range "qwertyuiopasdfghjklzxcvbnm" {
		data := string(data)
		file, err := os.Open("../../static/" + data + ".json")
		if err != nil {
			t.Error(err)
		}
		defer file.Close()
		byteValue, _ := ioutil.ReadAll(file)
		err = pg.AddSequence(data, byteValue)

		if err != nil {
			log.Fatal(err)
		}
	}
}
