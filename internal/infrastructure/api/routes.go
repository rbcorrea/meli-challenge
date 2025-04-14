package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rbcorrea/meli-challenge/internal/application/usecase"
	"github.com/rbcorrea/meli-challenge/internal/infrastructure/api/handler"
)

func RegisterRoutes(
	app *fiber.App,
	shortenUseCase *usecase.ShortenURLUseCase,
	listUseCase *usecase.ListURLsUseCase,
	searchByCodeUseCase *usecase.SearchByCodeUseCase,
	deleteUseCase *usecase.DeleteURLUseCase,
) {
	app.Get("/health", handler.HealthCheck())

	app.Post("/", handler.ShortenURL(shortenUseCase))
	app.Get("/", handler.ListURLs(listUseCase))
	app.Get("/:code", handler.SearchByCode(searchByCodeUseCase))
	app.Delete("/:code", handler.DeleteURL(deleteUseCase))
}
