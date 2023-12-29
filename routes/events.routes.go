package routes

import (
	"github.com/PrinceNarteh/go-events-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func EventsRoutes(app *fiber.App) {
	app.Get("/events", controllers.GetEvents)
	app.Post("/events", controllers.CreateEvent)
}
