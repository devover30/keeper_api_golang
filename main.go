package main

import (
	"context"
	"log"
	"os"

	"appstack.xyz/keeper_rest_api/controllers"
	"appstack.xyz/keeper_rest_api/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router *gin.Engine
	logger *log.Logger
	ctx    context.Context
)

func init() {
	logger = log.New(os.Stdout, "[auth-gateway] ", log.LstdFlags)
}

func main() {
	//logger.Println(os.Environ())

	err := godotenv.Load()
	if err != nil {
		logger.Panic("Error loading .env file")
	}

	router = gin.Default()

	v1Route := router.Group("/api/v1", middlewares.AuthTokenMiddleware(), middlewares.TokenValidateMiddleware())

	credntialController := controllers.NewCredentialController(logger, ctx)

	{
		v1Route.POST("/credentials", credntialController.CreateCredentialAction)
		v1Route.GET("/credentials", credntialController.FetchCredentialsAction)
		v1Route.DELETE("/credentials/:id", credntialController.DeleteCredentailAction)
		v1Route.PUT("/credentials/:id", credntialController.UpdateCredentailAction)
	}
	router.Run()
}
