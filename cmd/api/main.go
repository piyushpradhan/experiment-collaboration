package api

import (
	service "collaboration/services/api"
)

func InitializeApi() {
	apiService := service.NewApiService()

	allowedOrigins := []string{
		"https://experiment.piyushpradhan.space",
		"http://localhost",
		"https://showoff-frontend.vercel.app",
	}

	service.MakeHTTPHandler(apiService, allowedOrigins)
}
