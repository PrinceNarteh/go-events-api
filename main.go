package main

import (
	"log"

	"github.com/PrinceNarteh/go-events-api/controllers"
	"github.com/PrinceNarteh/go-events-api/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// initialize database
	db.InitDB()

	server := gin.Default()
	server.GET("/events", controllers.GetEvents)
	server.POST("/events", controllers.CreateEvent)
	server.Run(":4000")
}
