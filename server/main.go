package main

import (
	"log"
	"net/http"

	openapi "github.com/patriceandala/jumia/server/go"
)

func main() {
	log.Printf("Server started")

	DefaultApiService := openapi.NewDefaultApiService()
	DefaultApiController := openapi.NewDefaultApiController(DefaultApiService)

	router := openapi.NewRouter(DefaultApiController)

	if router != nil {
		log.Printf("Router is nil")
	}

	err := http.ListenAndServe(":8085", router)
	if err != nil {
		log.Printf("AN ERROR OCCURED %v", err)
		return
	}

}
