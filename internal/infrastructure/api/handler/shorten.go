package handler

import (
	"github.com/gofiber/fiber/v2"
)

type shortenedURLResponse struct {
	ShortenedURL string `json:"shortened_url"`
	OriginalURL  string `json:"original_url"`
}

func ShortenURL() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request struct {
			URL string `json:"url"`
		}
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		// // Call the use case to shorten the URL
		// shortenedURL, err := shortenUseCase.Shorten(request.URL)
		// if err != nil {
		// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to shorten URL"})
		// }

		return c.SendStatus(fiber.StatusOK).JSON(fiber.Map{
			"shortened_url": "http://short.url/" + request.URL,
		})
	}
}
