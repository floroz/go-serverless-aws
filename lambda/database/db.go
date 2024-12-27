package database

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"golang.org/x/crypto/bcrypt"
)

type DBClient struct {
	dynamoDB *dynamodb.DynamoDB
}

const (
	TableName      = "user_table"
	minPasswordLen = 4
	minUsernameLen = 3
)

// Pointer to singleton instance
var db *DBClient

func GetDBClient() *DBClient {
	dbSession := session.Must(session.NewSession())
	dynamo := dynamodb.New(dbSession)

	if db == nil {
		db = &DBClient{
			dynamoDB: dynamo,
		}
	}

	return db
}

func (u *DBClient) DoesUserExist(username string) (bool, error) {
	username = strings.TrimSpace(username)

	if db == nil {
		return false, fmt.Errorf("no db client yet")
	}

	result, err := u.dynamoDB.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TableName),
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

func (dbClient *DBClient) CreateUser(username, password string) error {
	if db == nil {
		return fmt.Errorf("no db client yet")
	}

	// sanitize
	trimmedUsername, trimmedPassword := strings.TrimSpace(username), strings.TrimSpace(password)

	// validate inputs
	if trimmedUsername == "" || len(trimmedUsername) < minUsernameLen {
		return fmt.Errorf("invalid username, at least 3 char long")
	}

	if trimmedPassword == "" || len(trimmedPassword) < minPasswordLen {
		return fmt.Errorf("invalid password, at least 4 char long")
	}

	// check if exists already
	if exist, err := dbClient.DoesUserExist(trimmedUsername); err != nil {
		return fmt.Errorf("something went wrong")
	} else if exist {
		return fmt.Errorf("invalid username")
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(trimmedPassword), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	// create item
	item := &dynamodb.PutItemInput{
		TableName: aws.String(TableName),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
			"password": {
				S: aws.String(string(hashedPassword)),
			},
		},
	}

	_, err = dbClient.dynamoDB.PutItem(item)

	// err creating
	if err != nil {
		return err
	}

	return nil
}
