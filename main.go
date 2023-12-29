package main

import (
	"github.com/PrinceNarteh/go-events-api/configs"
	"github.com/PrinceNarteh/go-events-api/controllers"
	"github.com/PrinceNarteh/go-events-api/mongorm"
	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.LoadEnv()

	// initialize database
	mongorm.InitDB()

	app := fiber.New()
	app.Get("/events", controllers.GetEvents)
	app.Post("/events", controllers.CreateEvent)
	app.Listen(":4000")
}
