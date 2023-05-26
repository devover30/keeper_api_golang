package controllers

import (
	"context"
	"log"
	"net/http"

	"appstack.xyz/keeper_rest_api/exceptions"
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

		ginContext.JSON(http.StatusOK, creds)
	}
}

func (controller *CredentialController) DeleteCredentailAction(ginContext *gin.Context) {
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

	credID := ginContext.Param("id")

	if credID == "" {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request.Credential ID missing"})
		return
	}

	var data interface{}

	data, _ = ginContext.Get("user")

	if user, ok := data.(*models.UserEntity); ok {

		service := services.NewCredentialService(db)

		err := service.RemoveCredential(user, credID)
		if err != nil {
			controller.logger.Println(err)
			if err == exceptions.ErrCredentialNotFound {
				ginContext.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			if err == exceptions.ErrorServer {
				ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ginContext.JSON(http.StatusNoContent, "")
	}
}

func (controller *CredentialController) UpdateCredentailAction(ginContext *gin.Context) {
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

	credID := ginContext.Param("id")

	if credID == "" {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request.Credential ID missing"})
		return
	}

	var credentialRequest models.CredentialRequestDTO
	if err := ginContext.ShouldBindJSON(&credentialRequest); err != nil {
		controller.logger.Println(err)
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input!"})
		return
	}

	var data interface{}

	data, _ = ginContext.Get("user")

	if user, ok := data.(*models.UserEntity); ok {

		service := services.NewCredentialService(db)

		updatedCred, err := service.ReformCredential(&credentialRequest, user, credID)
		if err != nil {
			controller.logger.Println(err)
			if err == exceptions.ErrCredentialNotFound {
				ginContext.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			if err == exceptions.ErrorServer {
				ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ginContext.JSON(http.StatusOK, updatedCred)
	}
}
