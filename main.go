package main

import (
	"github.com/PrinceNarteh/go-events-api/controllers"
	"github.com/PrinceNarteh/go-events-api/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", controllers.GetEvents)
	server.POST("/events", controllers.CreateEvent)
	server.Run(":4000")
}
