package repository

import (
	"context"

	"github.com/rbcorrea/meli-challenge/internal/domain/entity"
)

type ShortenURLRepository interface {
	Save(ctx context.Context, shortURL *entity.ShortURL) error
	FindByShortURL(ctx context.Context, shortURL string) (*entity.ShortURL, error)
	FindByOriginalURL(ctx context.Context, originalURL string) (*entity.ShortURL, error)
	FindByCode(ctx context.Context, code string) (*entity.ShortURL, error)
	Update(ctx context.Context, shortURL string, update interface{}) error
	IncrementAccessCount(ctx context.Context, code string) error
}
