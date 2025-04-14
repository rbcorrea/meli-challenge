package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rbcorrea/meli-challenge/internal/application/usecase"
)

func RedirectByCode(redirectUseCase *usecase.RedirectUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		code := c.Params("code")
		if code == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Code is required",
			})
		}

		shortURL, err := redirectUseCase.Execute(c.Context(), code)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "URL not found",
			})
		}

		if !shortURL.IsActive {
			return c.Status(fiber.StatusGone).JSON(fiber.Map{
				"error": "URL has been deleted",
			})
		}

		if err := redirectUseCase.IncrementAccessCount(c.Context(), code); err != nil {
			c.Context().Logger().Printf("Failed to increment access count: %v", err)
		}

		return c.Redirect(shortURL.OriginalURL, fiber.StatusMovedPermanently)
	}
}
