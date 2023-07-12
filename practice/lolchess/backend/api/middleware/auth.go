package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const BearerSchema = "Bearer"
const AuthHeaderKey = "Authorization"

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(AuthHeaderKey)
		if token == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		c.Next()
	}
}
