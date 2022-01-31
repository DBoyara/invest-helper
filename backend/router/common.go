package router

import (
	"github.com/gofiber/fiber/v2"
)

// HealthCheck returns ok
func HealthCheck(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
