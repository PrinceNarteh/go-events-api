package main

import (
	"github.com/PrinceNarteh/go-events-api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/events", controllers.GetEvents)
	server.POST("/events", controllers.CreateEvent)
	server.Run(":4000")
}
