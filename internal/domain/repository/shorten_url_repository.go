package repository

import (
	"context"

	"github.com/rbcorrea/meli-challenge/internal/domain/entity"
)

// ShortenURLRepository define a interface para operações de persistência de URLs encurtadas
type ShortenURLRepository interface {
	// Save salva uma URL encurtada no repositório
	Save(ctx context.Context, shortURL *entity.ShortURL) error
	// FindByShortURL busca uma URL encurtada pelo código
	FindByShortURL(ctx context.Context, shortURL string) (*entity.ShortURL, error)
	// FindByOriginalURL busca uma URL encurtada pela URL original
	FindByOriginalURL(ctx context.Context, originalURL string) (*entity.ShortURL, error)
	// FindByCode busca uma URL encurtada pelo código
	FindByCode(ctx context.Context, code string) (*entity.ShortURL, error)
	// Update atualiza uma URL encurtada
	Update(ctx context.Context, shortURL string, update interface{}) error
}
