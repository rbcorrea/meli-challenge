package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rbcorrea/meli-challenge/internal/application/usecase"
)

func SearchByCode(searchUseCase *usecase.SearchByCodeUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		code := c.Params("code")

		shortURL, err := searchUseCase.Execute(c.Context(), code)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(shortURL)
	}
}
