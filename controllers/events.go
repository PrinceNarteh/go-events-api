package controllers

import (
	"fmt"
	"net/http"

	"github.com/PrinceNarteh/go-events-api/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(ctx *gin.Context) {
	var events []models.Event
	ctx.JSON(http.StatusOK, events)
}

func GetAllEvents(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "All Events")
}

func CreateEvent(ctx *gin.Context) {
	var event models.Event
	if err := ctx.ShouldBindJSON(event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	fmt.Println(event)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created"})
}
