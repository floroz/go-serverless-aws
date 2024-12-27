package database

import (
	"fmt"
	"lambda-func/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"golang.org/x/crypto/bcrypt"
)

type DynamoDBClient struct {
	databaseStore *dynamodb.DynamoDB
}

const (
	tableName = "user_table"
)

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

func (u *DynamoDBClient) DoesUserExist(username string) (bool, error) {
	if db == nil {
		return false, fmt.Errorf("no db client yet")
	}

	result, err := u.databaseStore.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			username: {
				S: aws.String(username),
			},
		},
	})

	// err retrieving
	if err != nil {
		return false, err
	}

	// no user found
	if result.Item == nil {
		return false, nil
	}

	// user found
	return true, nil
}

func (u *DynamoDBClient) CreateUser(user *types.RegisterUser) error {
	if db == nil {
		return fmt.Errorf("no db client yet")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	item := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(user.Username),
			},
			"password": {
				S: aws.String(string(hashedPassword)),
			},
		},
	}

	_, err = u.databaseStore.PutItem(item)

	// err putting
	if err != nil {
		return err
	}

	return nil

}
