package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type LambdaEvent struct {
	Username string `json:"username"`
}

func HandleRequest(event LambdaEvent) error {
	if event.Username == "" {
		log.Println("Username empty", event)
		return fmt.Errorf("username cannot be empty")
	}
	return nil
}

func main() {
	defer log.Println("Lambda completed.")

	log.Println("Starting lambda...")
	lambda.Start(HandleRequest)
}
