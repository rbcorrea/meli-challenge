package handler

import (
	"github.com/gofiber/fiber/v2"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		response := HealthResponse{
			Status: "UP",
		}

		return c.Status(fiber.StatusOK).JSON(response)
	}
}
