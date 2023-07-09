package middlewares

import (
	"encoding/json"
	"net/http"
	"os"

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

		baseURL := os.Getenv("AUTH_SERVER")

		req, err := http.NewRequest("GET", baseURL+"/user", nil)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "The server encountered an internal error while processing this request."})
			return
		}

		req.Header.Add("Authorization", header.Authorization)

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "The server encountered an internal error while processing this request."})
			return
		}

		println(resp.StatusCode)
		if resp.StatusCode == 401 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			return
		}

		if resp.StatusCode == 500 {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "The server encountered an internal error while processing this request."})
			return
		}

		defer resp.Body.Close()

		var user *models.UserEntity

		if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "The server encountered an internal error while processing this request."})
			return
		}

		ctx.Set("user", user)

		ctx.Next()
	}
}
