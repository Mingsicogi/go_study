package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	env := os.Getenv("ENV")
	if "" == env {
		env = "local"
	}

	godotenvErr := godotenv.Load(".resources/properties/.env_" + env)
	if godotenvErr != nil {
		fmt.Println("No .env file found")
		os.Exit(1)
	}

	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_USER:", os.Getenv("DB_USER"))
	fmt.Println("DB_PASS:", os.Getenv("DB_PASS"))

	engine := gin.Default()
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	ginGonicErr := engine.Run()
	if ginGonicErr != nil {
		return
	}
}
