package sequenceRepo

import (
	"PP/worker/Math"
	"encoding/json"
	"log"
)

type IRepo interface {
	GetSequence(name string) (*Math.FloatSequence, error)
}

func GetSequenceJson(byteValue []byte) (*Math.FloatSequence, error) {

	var result map[string][]float64
	err := json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return Math.NewFloatSequence(result["data"]), nil
}
