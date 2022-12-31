package main

import (
	"github.com/go-ini/ini"
	"log"
	"os"
	"strconv"

	"PP/app/internal/generated/restapi"
	"PP/app/internal/generated/restapi/operations"
	"PP/app/internal/handlers"
	"github.com/go-openapi/loads"
)

func main() {
	if err := LoadAppConfig(); err != nil {
		log.Fatalln(err)
	}

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewAPIAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer func() {
		if err := server.Shutdown(); err != nil {
			log.Fatalln(err)
		}
	}()

	server.Port, _ = strconv.Atoi(os.Getenv("HTTP_PORT"))

	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(handlers.HealthCheckHandler)
	api.SayHelloHandler = operations.SayHelloHandlerFunc(handlers.SayHelloHandler)
	api.ProduceMessageHandler = operations.ProduceMessageHandlerFunc(handlers.ProduceMessageHandler)
	api.GetResultsHandler = operations.GetResultsHandlerFunc(handlers.GetResultsHandler)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func LoadAppConfig() error {
	cfg, err := ini.Load("config/app.ini")

	if err != nil {
		return err
	}

	err = os.Setenv("JOBS_SENT", "0")
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
