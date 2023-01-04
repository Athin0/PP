package sequenceRepo

import (
	"PP/worker/genericMath"
	"encoding/json"
	"log"
)

type IRepo interface {
	GetSequence(name string) (*genericMath.FloatSequence, error)
}

func GetSequenceJson(byteValue []byte) (*genericMath.FloatSequence, error) {

	var result map[string][]float64
	err := json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return genericMath.NewFloatSequence(result["data"]), nil
}
