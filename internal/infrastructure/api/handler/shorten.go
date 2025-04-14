package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rbcorrea/meli-challenge/internal/application/dto"
	"github.com/rbcorrea/meli-challenge/internal/application/usecase"
)

func ShortenURL(shortenUseCase *usecase.ShortenURLUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request dto.ShortenURLRequest
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		response, err := shortenUseCase.Execute(c.Context(), &request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(response)
	}
}
