package controllers

import (
	"context"
	"log"
	"net/http"

	"appstack.xyz/keeper_rest_api/lib"
	"appstack.xyz/keeper_rest_api/models"
	"appstack.xyz/keeper_rest_api/services"
	"github.com/gin-gonic/gin"
)

type CredentialController struct {
	logger  *log.Logger
	context context.Context
}

func NewCredentialController(l *log.Logger, ctx context.Context) *CredentialController {
	return &CredentialController{logger: l, context: ctx}
}

func (controller *CredentialController) CreateCredentialAction(ginContext *gin.Context) {
	db, err := lib.ConnectToDB()
	if err != nil {
		controller.logger.Println(err)
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "The server encountered an internal error while processing this request."})
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		controller.logger.Println(err)
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "The server encountered an internal error while processing this request."})
		return
	}

	var credentialResquest models.CredentialRequestDTO
	if err := ginContext.ShouldBindJSON(&credentialResquest); err != nil {
		controller.logger.Println(err)
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input!"})
		return
	}

	service := services.NewCredentialService(db)

	var data interface{}

	data, _ = ginContext.Get("user")

	if user, ok := data.(*models.UserEntity); ok {

		newCred, err := service.PersistCredential(&credentialResquest, user)
		if err != nil {
			controller.logger.Println(err)
			ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ginContext.JSON(http.StatusCreated, newCred)
	}
}

func (controller *CredentialController) FetchCredentialsAction(ginContext *gin.Context) {
	db, err := lib.ConnectToDB()
	if err != nil {
		controller.logger.Println(err)
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "The server encountered an internal error while processing this request."})
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		controller.logger.Println(err)
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "The server encountered an internal error while processing this request."})
		return
	}

	var data interface{}

	data, _ = ginContext.Get("user")

	if user, ok := data.(*models.UserEntity); ok {

		service := services.NewCredentialService(db)

		creds, err := service.AcquireCredentials(user)
		if err != nil {
			controller.logger.Println(err)
			ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ginContext.JSON(http.StatusCreated, creds)
	}
}
