package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("unable to load .env file")
	}
	
	router := gin.Default()

	router.Run(":" + os.Getenv("PORT"))
}
