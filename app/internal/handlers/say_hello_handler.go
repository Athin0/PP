package handlers

import (
	"fmt"
	"log"

	"github.com/go-openapi/runtime/middleware"

	"PP/app/internal/generated/restapi/operations"
)

func SayHelloHandler(params operations.SayHelloParams) middleware.Responder {
	log.Printf("Hit GET /hello/{user} from %s\n", params.HTTPRequest.UserAgent())

	if params.User == "John" {
		return operations.NewSayHelloBadRequest().WithPayload("Fuck you John")
	}

	return operations.NewSayHelloOK().WithPayload(fmt.Sprintf("Hello, %s!", params.User))
}
