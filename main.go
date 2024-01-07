package main

import (
	"github.com/PrinceNarteh/go-events-api/configs"
	"github.com/PrinceNarteh/go-events-api/mongorm"
	"github.com/PrinceNarteh/go-events-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.InitEnvConfigs()

	// initialize database
	mongorm.InitDB()

	// Fiber instance
	app := fiber.New()

	// Routes
	routes.EventsRoutes(app)

	// Listening to Port
	app.Listen(":4000")
}
