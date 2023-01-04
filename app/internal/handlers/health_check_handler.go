package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-openapi/runtime/middleware"

	"PP/app/internal/generated/restapi/operations"
)

func HealthCheckHandler(params operations.CheckHealthParams) middleware.Responder {
	log.Printf("Hit GET /health/check from %s\n", params.HTTPRequest.UserAgent())

	kafkaVer, kafkaStat := CheckHealthKafka()

	log.Println("Application is healthy")

	workers, err := CheckHealthWorkers()
	if err != nil {
		log.Printf("Consumer error (app readMessage): %v \n", err)
		return operations.NewCheckHealthInternalServerError().WithPayload("Internal server error: " + err.Error()) // may be better return more data
	}
	return operations.NewCheckHealthOK().WithPayload(
		&operations.CheckHealthOKBody{
			App: &operations.CheckHealthOKBodyApp{
				Version: os.Getenv("APPLICATION_VERSION"),
				Status:  true,
			},
			Kafka: &operations.CheckHealthOKBodyKafka{
				Version: kafkaVer,
				Status:  kafkaStat,
			},
			Workers: workers,
		},
	)
}

func CheckHealthKafka() (string, bool) {
	cli, err := kafka.NewAdminClient(
		&kafka.ConfigMap{
			"metadata.broker.list": fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT")),
		},
	)
	if err != nil {
		log.Println("Unable to connect to kafka")

		return "-", false
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := cli.ClusterID(ctx)
	if err != nil {
		log.Println("Unable to connect to kafka")

		return "-", false
	}

	log.Println("Kafka is healthy")

	return id, true
}

func CheckHealthWorkers() ([]*operations.CheckHealthOKBodyWorkersItems0, error) {
	partitions, errPart := strconv.Atoi(os.Getenv("PARTITIONS"))
	if errPart != nil {
		log.Printf("err in cheackHealthWorkers get PARTITIONS %v", errPart)
		return []*operations.CheckHealthOKBodyWorkersItems0{}, errPart
	}

	c, err := kafka.NewConsumer(
		&kafka.ConfigMap{
			"bootstrap.servers": fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT")),
			"group.id":          "healthAppGroup",
			"auto.offset.reset": "earliest",
		},
	)
	if err != nil {
		log.Println(err)
		return []*operations.CheckHealthOKBodyWorkersItems0{}, err
	}

	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"metadata.broker.list": fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT")),
		},
	)
	if err != nil {
		log.Println(err)

		return []*operations.CheckHealthOKBodyWorkersItems0{}, err
	}

	defer func() { _ = c.Close() }()
	defer p.Close()

	err = c.SubscribeTopics([]string{os.Getenv("KAFKA_TOPIC_WORKER_HEALTH_READ")}, nil)
	if err != nil {
		log.Println(err)

		return []*operations.CheckHealthOKBodyWorkersItems0{}, err
	}

	topic := os.Getenv("KAFKA_TOPIC_WORKER_HEALTH_WRITE")
	for i := 0; i < partitions; i++ {
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: int32(i),
			},
			Value: []byte("true"),
		}, nil)
	}
	if err != nil {
		log.Printf("Error in produce message to kafka: %v \n", err)
		return []*operations.CheckHealthOKBodyWorkersItems0{}, err
	}

	workers := make(map[string]bool)
	result := make([]*operations.CheckHealthOKBodyWorkersItems0, 0)

	for i := 0; i < partitions; i++ {
		fmt.Println("Here 7:", i)
		msg, err := c.ReadMessage(10 * time.Second)
		if err != nil {
			log.Printf("Consumer error (app readMessage): %v \n", err)
			return []*operations.CheckHealthOKBodyWorkersItems0{}, err
		}
		data := operations.CheckHealthOKBodyWorkersItems0{}
		err = json.Unmarshal(msg.Value, &data)
		if err != nil {
			log.Println(err)
			return []*operations.CheckHealthOKBodyWorkersItems0{}, err
		}
		if _, ok := workers[data.UUID]; !ok {
			result = append(result, &data)
			workers[data.UUID] = true
		}
	}
	return result, nil
}
