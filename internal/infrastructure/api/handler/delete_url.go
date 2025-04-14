package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rbcorrea/meli-challenge/internal/application/dto"
	"github.com/rbcorrea/meli-challenge/internal/application/usecase"
)

func DeleteURL(deleteUseCase *usecase.DeleteURLUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		code := c.Params("code")
		request := &dto.DeleteURLRequest{
			Code: code,
		}

		response, err := deleteUseCase.Execute(c.Context(), request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(response)
	}
}
