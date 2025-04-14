package usecase

import (
	"context"

	"github.com/rbcorrea/meli-challenge/internal/domain/entity"
)

type ShortenURLUseCaseInterface interface {
	Execute(ctx context.Context, originalURL string) (*entity.ShortURL, error)
}
