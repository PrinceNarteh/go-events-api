package main

import (
	"net/http"

	"github.com/PrinceNarteh/go-events-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)

	server.Run(":4000")
}

func getEvents(ctx *gin.Context) {
	events := models.GetAllEvents()
	ctx.JSON(http.StatusOK, events)
}
