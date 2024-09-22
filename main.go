package main

import (
	"log"
	"net/http"

	"microservice-apis-go/oapi_codegen/generated/api"
	"microservice-apis-go/oapi_codegen/server"

	"github.com/gin-gonic/gin"
)

func main() {
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := server.NewServer()

	r := gin.Default()

	api.RegisterHandlers(r, server)

	// And we serve HTTP until the world ends.

	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:80",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
