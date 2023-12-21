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
