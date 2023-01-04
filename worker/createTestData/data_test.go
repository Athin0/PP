package createTestData

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
)

const arr = "qwertyuiopasdfghjklzxcvbnm"

type Message struct {
	Data []float64 `json:"data"`
}

func TestGenerete(t *testing.T) {
	for _, a := range arr {
		CreateMake(string(a))
	}
}
func Gnerate() []float64 {
	var i int
	arr := make([]float64, 100)
	for i < 100 {
		arr[i] = rand.Float64()
		i++
	}
	return arr
}

func CreateMake(text string) {
	f, _ := os.Create(text + ".json")
	defer f.Close()
	m := Message{Gnerate()}

	b, err := json.Marshal(m)
	fmt.Println(string(b))
	if err != nil {
		log.Println(err)
	}
	_, err = f.Write(b)
	if err != nil {
		log.Println(err)
	}
}
