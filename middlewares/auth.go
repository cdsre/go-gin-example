package middlewares

import (
	"github.com/gin-gonic/gin"
	"my-rest-api/utils"
	"net/http"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not Authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not Authorized"})
		return
	}

	c.Set("userId", userId)
	c.Next()
}
