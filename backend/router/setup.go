package router

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setups all the Routes
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	common := api.Group("/health")
	common.Get("", HealthCheck)

	PASSGEN = api.Group("/pass-gen")
	SetupPassGenRoutes()

	TRADING = api.Group("/trading")
	SetupTradingRoutes()

	FUTURES = api.Group("/futures")
	SetupFuturesRoutes()
}
