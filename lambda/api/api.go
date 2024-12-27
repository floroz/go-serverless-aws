package api

import "lambda-func/database"

type ApiClient struct {
	dbStore *database.DynamoDBClient
}

func (c *ApiClient) CreateUser(usrname, password string) error {
	return nil
}

// Pointer to singleton instance
var apiClient *ApiClient

func GetApiClient() *ApiClient {
	if apiClient == nil {
		db := database.GetDynamoDBClient()
		apiClient = &ApiClient{
			dbStore: db,
		}
	}

	return apiClient
}
