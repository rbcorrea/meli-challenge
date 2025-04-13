package api

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	// Health Check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/shorten", handler.ShortenURL(shortenUseCase))
	app.Get("/:code", handler.ResolveShortURL(resolveUseCase))
	app.Delete("/:code", handler.DeleteURL(deleteUseCase))
}
