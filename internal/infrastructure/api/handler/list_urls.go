package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rbcorrea/meli-challenge/internal/application/dto"
	"github.com/rbcorrea/meli-challenge/internal/application/usecase"
)

func ListURLs(listUseCase *usecase.ListURLsUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := &dto.ListURLsRequest{
			Page:     1,
			PageSize: 10,
		}

		if err := c.QueryParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid query parameters",
			})
		}

		if request.Page < 1 {
			request.Page = 1
		}
		if request.PageSize < 1 {
			request.PageSize = 10
		}
		if request.PageSize > 100 {
			request.PageSize = 100
		}

		response, err := listUseCase.Execute(c.Context(), request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(response)
	}
}
