package cmd

import (
	"backend.com/go-backend/app/config"
	"backend.com/go-backend/app/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Server() *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	config.ConnectDatabase()
	router := routers.SetupRouter(config.SessionStorage())

	return router
}
