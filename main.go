package main

import (
	"context"
	"log"
	"os"

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

	//v1Route := router.Group("/api/v1")

	//memberController := controllers.NewMemberController(logger, ctx)
	//authController := controllers.NewAuthController(logger, ctx)
	////router.GET("/members/:mobile", memberController.GetByMobileAction)
	//{
	//	v1Route.POST("/members", memberController.CreateMemberAction)
	//	v1Route.POST("/members/verify", memberController.VerifyAction)
	//	v1Route.POST("/authentication", authController.LoginMemberAction)
	//	v1Route.POST("/authentication/verify", authController.VerifyMemberAction)
	//	v1Route.GET("/authentication/verify", authController.VerifyTokenAction)
	//	v1Route.POST("/authentication/refresh", authController.RefreshTokenAction)
	//}
	router.Run()
}
