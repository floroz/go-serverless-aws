package main

import (
	"fmt"
	"lambda-func/api"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type LambdaEvent struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleRequest(event LambdaEvent) error {
	if event.Username == "" {
		log.Println("Username empty", event)
		return fmt.Errorf("username cannot be empty")
	}
	apiClient := api.GetApiClient()

	return apiClient.CreateUser(event.Username, event.Password)
}

func main() {
	defer log.Println("Lambda completed.")

	log.Println("Starting lambda...")
	lambda.Start(HandleRequest)
}
