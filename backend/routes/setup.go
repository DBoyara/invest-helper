package routes

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setups all the Routes
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	common := api.Group("/health")
	common.Get("", HealthCheck)

}
