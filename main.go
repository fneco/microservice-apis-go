package main

import (
	"log"
	"net/http"

	"microservice-apis-go/oapi_codegen/generated/api"
	"microservice-apis-go/oapi_codegen/server"

	middleware "github.com/oapi-codegen/gin-middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	server := server.NewServer()

	validator, err := middleware.OapiValidatorFromYamlFile("./oas.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}

	r := gin.Default()
	r.Use(validator)

	api.RegisterHandlers(r, server)

	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:80",
	}

	log.Fatal(s.ListenAndServe())
}
