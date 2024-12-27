package database

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBClient struct {
	databaseStore *dynamodb.DynamoDB
}

// Pointer to singleton instance
var db *DynamoDBClient

func GetDynamoDBClient() *DynamoDBClient {
	dbSession := session.Must(session.NewSession())
	dynamo := dynamodb.New(dbSession)

	if db == nil {
		db = &DynamoDBClient{
			databaseStore: dynamo,
		}
	}

	return db
}
