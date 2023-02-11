package MemoryRepo

import (
	"PP/worker/Math"
	"PP/worker/sequenceRepo"
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type MemoryRepo struct {
}

func (m MemoryRepo) GetSequence(name string) (*Math.FloatSequence, error) {
	file, err := os.Open("../static/" + name + ".txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	seq, err := m.GetSequenceTxt(file)
	//seq, err = m.GetSequenceJson(file)

	return seq, nil

}

func (MemoryRepo) GetSequenceJson(file *os.File) (*Math.FloatSequence, error) {
	byteValue, _ := ioutil.ReadAll(file)
	return sequenceRepo.GetSequenceJson(byteValue)
}

func (MemoryRepo) GetSequenceTxt(file *os.File) (*Math.FloatSequence, error) {
	seq := Math.FloatSequence{}

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

func (MemoryRepo) GetSequenceJson2(file *os.File) (*Math.FloatSequence, error) {
	seq := Math.FloatSequence{}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Println(err)
		}
		if err == io.EOF {
			break
		}
		data := strings.Split(line[:len(line)-2], " ")

		for _, snum := range data {
			f, err := strconv.ParseFloat(snum, 64)
			if err != nil {
				log.Println(err)
			}
			seq.Append(f)
		}

		if err != nil {
			log.Println(err)
		}
	}
	return &seq, nil
}
