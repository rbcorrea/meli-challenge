package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rbcorrea/meli-challenge/internal/application/usecase"
)

type App struct {
	fiber *fiber.App
}

func NewApp(
	shortenUseCase *usecase.ShortenURLUseCase,
	listUseCase *usecase.ListURLsUseCase,
	searchByCodeUseCase *usecase.SearchByCodeUseCase,
	deleteUseCase *usecase.DeleteURLUseCase,
) *App {
	app := fiber.New(fiber.Config{
		AppName: "URL Shortener API",
	})

	app.Use(recover.New())
	app.Use(logger.New())

	RegisterRoutes(app, shortenUseCase, listUseCase, searchByCodeUseCase, deleteUseCase)

	return &App{
		fiber: app,
	}
}

func (a *App) Listen(addr string) error {
	return a.fiber.Listen(addr)
}
