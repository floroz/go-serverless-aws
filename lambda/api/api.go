package api

import (
	"fmt"
	"lambda-func/database"
	"log"
)

type ApiClient struct {
	userStore database.UserStore
}

// Pointer to singleton instance
var apiClient *ApiClient

func GetApiClient() *ApiClient {
	if apiClient == nil {
		db := database.GetDBClient()
		apiClient = &ApiClient{
			userStore: db,
		}
	}

	return apiClient
}

func (a *ApiClient) CreateUser(usr, pwd string) error {

	err := a.userStore.CreateUser(usr, pwd)

	if err != nil {
		log.Printf("Failed to create user.\n")
		return fmt.Errorf("error while creating the user, error: %w", err)
	}

	log.Printf("User %s created\n", usr)

	return nil

}
