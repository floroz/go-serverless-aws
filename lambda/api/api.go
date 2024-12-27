package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiClient struct {
	dbStore *database.DBClient
}

func (c *ApiClient) CreateUser(usrname, password string) error {
	return nil
}

// Pointer to singleton instance
var apiClient *ApiClient

func GetApiClient() *ApiClient {
	if apiClient == nil {
		db := database.GetDBClient()
		apiClient = &ApiClient{
			dbStore: db,
		}
	}

	return apiClient
}

func (a *ApiClient) RegisterUserHandler(usr types.RegisterUser) error {

	err := a.dbStore.CreateUser(usr.Username, usr.Password)

	if err != nil {
		return fmt.Errorf("error while creating the user, error: %w", err)
	}

	return nil

}
