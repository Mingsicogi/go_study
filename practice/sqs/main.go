package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/welcome", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	r.Run()
}
