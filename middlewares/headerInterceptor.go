package middlewares

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"appstack.xyz/keeper_rest_api/models"
	"github.com/gin-gonic/gin"
)

type Header struct {
	Authorization string `header:"Authorization" binding:"required"`
}

var header = &Header{}

func AuthTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := c.ShouldBindHeader(header); err != nil {
			c.JSON(400, gin.H{"error": "Invalid Request.Authorization Header not present."})
			return
		}

		c.Next()

	}
}

func TokenValidateMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := strings.Split(header.Authorization, "Bearer ")[1]
		fmt.Println(token)

		baseURL := os.Getenv("AUTH_SERVER")

		req, err := http.NewRequest("GET", baseURL+"/authentication/verify", nil)

		req.Header.Add("Authorization", header.Authorization)

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			log.Fatal("ooopsss an error occurred, please try again")
		}

		println(resp.StatusCode)

		defer resp.Body.Close()

		var user *models.UserEntity

		if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
			log.Panicln("Request Error..Error reading body to json")
		}

		ctx.Set("user", user)

		ctx.Next()
	}
}
