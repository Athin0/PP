package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-openapi/runtime/middleware"
	"log"
	"os"
	"strconv"

	"PP/app/internal/generated/restapi/operations"
)

func GetResultsHandler(params operations.GetResultsParams) middleware.Responder {
	log.Printf("Hit GET /get/results from %s\n", params.HTTPRequest.UserAgent())

	jobs, _ := strconv.Atoi(os.Getenv("JOBS_SENT"))
	log.Printf("Active jobs to get: %d\n", jobs)
	if jobs != 0 {
		c, err := kafka.NewConsumer(
			&kafka.ConfigMap{
				"bootstrap.servers": fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT")),
				"group.id":          "resultsGroup",
				"auto.offset.reset": "earliest",
			},
		)
		if err != nil {
			log.Println(err)

			return operations.NewGetResultsInternalServerError().WithPayload("Internal server error")
		}
		defer func() { _ = c.Close() }()

		err = c.SubscribeTopics([]string{os.Getenv("KAFKA_TOPIC_READ")}, nil)
		if err != nil {
			log.Fatalln(err)
		}

		result := make([]*operations.GetResultsOKBodyItems0, 0)

		for i := 0; i < jobs; i++ {
			//This is a convenience API that wraps Poll() and only returns messages or errors.
			//All other event types are discarded.
			//todo may be better add a Pool
			msg, err := c.ReadMessage(-1)
			if err != nil {
				log.Printf("Consumer error: %v (%v)\n", err, msg)
				break
			}

			data := operations.GetResultsOKBodyItems0{}
			err = json.Unmarshal(msg.Value, &data)
			if err != nil {
				log.Printf("Unmarshall of msg error: %v \n", err)
				break
			}
			result = append(result, &data)
		}

		err = os.Setenv("JOBS_SENT", "0")
		if err != nil {
			log.Println(err)

			return operations.NewGetResultsInternalServerError().WithPayload("Internal server error")
		}

		return operations.NewGetResultsOK().WithPayload(result)
	}

	return operations.NewGetResultsOK().WithPayload(make([]*operations.GetResultsOKBodyItems0, 0))
}
