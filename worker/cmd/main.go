package main

import (
	"PP/worker/sequenceRepo"
	"PP/worker/sequenceRepo/PgRepo"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-ini/ini"
	uuid "github.com/satori/go.uuid"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"PP/worker/Math"
	"PP/worker/astParser"
	"PP/worker/asyncDispatching"
	"PP/worker/grammar/lexer"
	"PP/worker/grammar/parser"
)

type JobRequest struct {
	Data string `json:"data"`
	UUID string `json:"uuid"`
}

type HealthCheck struct {
	UUID   string `json:"uuid"`
	Status bool   `json:"status"`
}

var ErrParsingString = errors.New("error parsing string")

func main() {
	log.Println("Worker started!")

	if err := LoadWorkerConfig(); err != nil {
		log.Fatalln(err)
	}
	db, err := PgRepo.InitDB()
	if err != nil {
		log.Println("error of init db: ", err)
		return
	}
	go HealthCheckProcess()

	c, err := kafka.NewConsumer(
		&kafka.ConfigMap{
			"bootstrap.servers": fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT")),
			"group.id":          "workerGroup",
			"auto.offset.reset": "latest",
		},
	)
	if err != nil {
		log.Fatalln(err)
	}

	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"metadata.broker.list": fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT")),
		},
	)
	if err != nil {
		log.Fatalln(err)
	}

	defer func() { _ = c.Close() }()
	defer p.Close()

	err = c.SubscribeTopics([]string{os.Getenv("KAFKA_TOPIC_READ")}, nil)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			log.Printf("Consumer error: %v (%v)\n", err, msg)

			continue
		}
		log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))

		request := JobRequest{}
		err = json.Unmarshal(msg.Value, &request)
		if err != nil {
			log.Println(err)
			continue
		}

		var bytes []byte
		result, err := CalculateSequence(request.Data, db)
		if err != nil {
			log.Println(err)
			bytes, _ = json.Marshal(
				map[string]interface{}{
					"uuid":  request.UUID,
					"error": err.Error(),
				},
			)
		} else {
			bytes, _ = json.Marshal(
				map[string]interface{}{
					"result": result,
					"uuid":   request.UUID,
				},
			)
			log.Printf("Calculation result: %f\n", result)
		}

		topic := os.Getenv("KAFKA_TOPIC_WRITE")
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: bytes,
		}, nil)

	}
}

func InitTestingSequence(path string) (*Math.FloatSequence, error) {
	seq := Math.FloatSequence{}

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

func CalculateSequence(text string, repo sequenceRepo.IRepo) (float64, error) {
	lex := lexer.New([]rune(text))

	if _bsr, errs := parser.Parse(lex); len(errs) != 0 {
		log.Printf("Err in Parser: %v", ErrParsingString)

		return 0, ErrParsingString
	} else {
		root, err := astParser.BuildAST(_bsr.GetRoot(), nil, repo)
		if err != nil {
			return 0, err
		}
		disp := asyncDispatching.NewDispatcher(root)

		return asyncDispatching.Traverse(disp), nil
	}
}

func HealthCheckProcess() {
	c, err := kafka.NewConsumer(
		&kafka.ConfigMap{
			"bootstrap.servers": fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT")),
			"group.id":          "healthWorkerGroup",
			"auto.offset.reset": "latest",
		},
	)
	if err != nil {
		log.Fatalln(err)
	}

	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"metadata.broker.list": fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT")),
		},
	)
	if err != nil {
		log.Fatalln(err)
	}

	defer func() { _ = c.Close() }()
	defer p.Close()

	err = c.SubscribeTopics([]string{os.Getenv("KAFKA_TOPIC_WORKER_HEALTH_READ")}, nil)
	if err != nil {
		log.Fatalln(err)
	}

	health, _ := json.Marshal(
		HealthCheck{
			UUID:   os.Getenv("WORKER_UUID"),
			Status: true,
		},
	)

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			log.Printf("Consumer error: %v (%v)\n", err, msg)

			continue
		}

		log.Printf("Received health check on %s\n", msg.TopicPartition)

		topic := os.Getenv("KAFKA_TOPIC_WORKER_HEALTH_WRITE")
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: health,
		}, nil)
		if err != nil {
			log.Printf("Producer error: %v (%v)\n", err, msg)

			continue
		}

		log.Printf("Produced health check on %s\n", msg.TopicPartition)
	}
}

func LoadWorkerConfig() error {
	cfg, err := ini.Load("config/worker.ini")
	if err != nil {
		return err
	}

	err = os.Setenv("WORKER_UUID", uuid.NewV4().String())
	if err != nil {
		return err
	}

	if os.Getenv("DOCKER_COMPOSE") == "true" {
		section := cfg.Section("compose")
		for _, key := range section.Keys() {
			err = os.Setenv(key.Name(), key.Value())
			if err != nil {
				return err
			}
		}

		return nil
	} else {
		section := cfg.Section("debug")
		for _, key := range section.Keys() {
			err = os.Setenv(key.Name(), key.Value())
			if err != nil {
				return err
			}
		}

		return nil
	}
}
