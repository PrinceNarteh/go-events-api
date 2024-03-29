package controllers

import (
	"net/http"
	"time"

	"github.com/PrinceNarteh/go-events-api/models"
	"github.com/PrinceNarteh/go-events-api/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var eventsCollection = "events"

func GetEvents(ctx *fiber.Ctx) error {
	var events models.Event
	doc, err := events.Find(ctx.Context(), eventsCollection, bson.M{"name": "Xmas Beach"})
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": doc})
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

	event.Create(ctx.Context(), eventsCollection, &event)

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": event})
}
