package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/PrinceNarteh/go-events-api/models"
	"github.com/PrinceNarteh/go-events-api/utils"
	"github.com/gofiber/fiber/v2"
)

var eventsCollection = "events"

func GetEvents(ctx *fiber.Ctx) error {
	var events []models.Event
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": events})
}

func GetAllEvents(ctx *fiber.Ctx) error {
	return ctx.JSON(http.StatusOK, "All Events")
}

func CreateEvent(ctx *fiber.Ctx) error {
	var event models.Event
	if err := ctx.BodyParser(&event); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	event.DateTime = time.Now()

	if errs := utils.ValidateStruct(&event); len(errs) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": errs})
	}

	event.Create(context.Background(), eventsCollection, &event)

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": event})
}
