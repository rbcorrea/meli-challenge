package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rbcorrea/meli-challenge/internal/application/usecase"
	"github.com/rbcorrea/meli-challenge/internal/infrastructure/api/handler"
)

func RegisterRoutes(
	app *fiber.App,
	shortenUseCase *usecase.ShortenURLUseCase,
	searchByCodeUseCase *usecase.SearchByCodeUseCase,
	redirectUseCase *usecase.RedirectUseCase,
	deleteUseCase *usecase.DeleteURLUseCase,
) {
	app.Get("/health", handler.HealthCheck())

	app.Post("/", handler.ShortenURL(shortenUseCase))
	app.Get("/stats/:code", handler.SearchByCode(searchByCodeUseCase))
	app.Delete("/:code", handler.DeleteURL(deleteUseCase))
	app.Get("/:code", handler.RedirectByCode(redirectUseCase))
}
