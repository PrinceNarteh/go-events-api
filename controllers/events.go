package controllers

import (
	"net/http"

	"github.com/PrinceNarteh/go-events-api/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(ctx *gin.Context) {
	events := models.GetAllEvents()
	ctx.JSON(http.StatusOK, events)
}

func GetAllEvents(ctx *gin.Context) {
	events := models.GetAllEvents()
	ctx.JSON(http.StatusOK, events)
}

func CreateEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	event.ID = 1
	event.UserID = 1
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created"})
}
