package usecase

import (
	"context"
	"fmt"

	"github.com/rbcorrea/meli-challenge/internal/domain/entity"
	"github.com/rbcorrea/meli-challenge/internal/domain/repository"
	"github.com/redis/go-redis/v9"
)

type RedirectUseCase struct {
	redis      *redis.Client
	repository repository.ShortenURLRepository
}

func NewRedirectUseCase(redis *redis.Client, repository repository.ShortenURLRepository) *RedirectUseCase {
	return &RedirectUseCase{
		redis:      redis,
		repository: repository,
	}
}

func (u *RedirectUseCase) Execute(ctx context.Context, code string) (*entity.ShortURL, error) {
	// Tenta buscar no Redis primeiro
	originalURL, err := u.redis.Get(ctx, code).Result()
	if err == nil {
		return &entity.ShortURL{
			OriginalURL: originalURL,
			ShortURL:    "https://me.li/" + code,
			Code:        code,
			IsActive:    true,
		}, nil
	}

	// Se não encontrar no Redis, busca no MongoDB
	url, err := u.repository.FindByCode(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to find URL: %w", err)
	}

	return url, nil
}

func (u *RedirectUseCase) IncrementAccessCount(ctx context.Context, code string) error {
	return u.repository.IncrementAccessCount(ctx, code)
}
