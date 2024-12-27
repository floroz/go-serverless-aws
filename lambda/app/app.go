package app

import "lambda-func/api"

type App struct {
	ApiClient *api.ApiClient
}

var app *App

func GetApp() {
	if app == nil {
		app = &App{
			ApiClient: api.GetApiClient(),
		}
	}
}
