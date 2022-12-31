package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"log"
	"os"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-openapi/runtime/middleware"

	"PP/app/internal/generated/restapi/operations"
	"PP/app/internal/mthprsr"
)

type JobMessage struct {
	operations.ProduceMessageBody
	UUID string `json:"uuid"`
}

func ProduceMessageHandler(params operations.ProduceMessageParams) middleware.Responder {
	log.Printf("Hit POST /produce from %s\n", params.HTTPRequest.UserAgent())

	job := JobMessage{
		ProduceMessageBody: params.Request,
		UUID:               uuid.NewV4().String(),
	}

	bytes, err := json.Marshal(job)
	if err != nil {
		log.Println(err)

		return operations.NewProduceMessageInternalServerError().WithPayload("Internal server error")
	}

	/*
		p, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": "host1:9092,host2:9092",
			"client.id": socket.gethostname(),
			"acks": "all"})

		if err != nil {
			fmt.Printf("Failed to create producer: %s\n", err)
			os.Exit(1)
		}


	*/
	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"metadata.broker.list": fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT")),
		},
	)
	if err != nil {
		log.Println(err)

		return operations.NewProduceMessageInternalServerError().WithPayload("Internal server error")
	}
	//internal function for validating output message correctness
	if !mthprsr.ValidateString(*params.Request.Data) {
		return operations.NewProduceMessageBadRequest().WithPayload("Bad request: unable to read string")
	}

	defer p.Close()

	//handle message delivery reports and possibly other event types (errors, stats, etc) concurrently
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					log.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	topic := os.Getenv("KAFKA_TOPIC_WRITE")
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: bytes,
	}, nil) // can be add chan to listen for the result of the send
	if err != nil {
		log.Println(err)

		return operations.NewProduceMessageInternalServerError().WithPayload("Internal server error")
	}

	jobs, _ := strconv.Atoi(os.Getenv("JOBS_SENT"))
	log.Printf("Jobs: %d\n", jobs)
	err = os.Setenv("JOBS_SENT", strconv.Itoa(jobs+1))
	if err != nil {
		log.Println(err)

		return operations.NewProduceMessageInternalServerError().WithPayload("Internal server error")
	}

	numUnflushed := p.Flush(15 * 1000)
	if numUnflushed != 0 {
		log.Printf("The number of outstanding events still un-flushed: %d", numUnflushed)
	}

	return operations.NewProduceMessageOK().WithPayload(fmt.Sprintf("Successfully produced message with uuid %s", job.UUID))
}
