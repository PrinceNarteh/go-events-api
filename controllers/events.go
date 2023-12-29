package controllers

import (
	"fmt"
	"net/http"

	"github.com/PrinceNarteh/go-events-api/models"
	"github.com/gofiber/fiber/v2"
)

func GetEvents(ctx *fiber.Ctx) error {
	var events []models.Event
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": events})
}

func GetAllEvents(ctx *fiber.Ctx) error {
	return ctx.JSON(http.StatusOK, "All Events")
}

func CreateEvent(ctx *fiber.Ctx) error {
	var event models.Event
	if err := ctx.BodyParser(event); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err})

	}
	fmt.Println(event)
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": event})
}
