package controllers

import (
	"fmt"
	"net/http"

	"github.com/PrinceNarteh/go-events-api/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

func GetEvents(ctx *fiber.Ctx) error {
	var events []models.Event
	return ctx.Status(http.StatusOK).JSON(fiber.Map{"message": events})
}

func GetAllEvents(ctx *fiber.Ctx) error {
	return ctx.JSON(http.StatusOK, "All Events")
}

func CreateEvent(ctx *fiber.Ctx) error {
	validate = validator.New(validator.WithRequiredStructEnabled())
	var event models.Event
	if err := ctx.BodyParser(&event); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err})

	}

	if ValidationErr := validate.Struct(&event); ValidationErr != nil {
		for _, err := range ValidationErr.(validator.ValidationErrors) {

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"message": ValidationErr})
	}

	fmt.Println(event)
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{"message": event})
}
