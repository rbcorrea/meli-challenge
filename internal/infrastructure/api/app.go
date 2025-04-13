package api

import (
	"github.com/gofiber/fiber/v2"
)

func NewApp() *fiber.App {

	app := fiber.New()

	// Middleware
	app.Use(func(c *fiber.Ctx) error {
		// Add any middleware you need here
		return c.Next()
	})

	RegisterRoutes(app)

	return app
}
