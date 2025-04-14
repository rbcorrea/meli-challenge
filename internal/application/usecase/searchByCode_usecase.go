package usecase

import (
	"context"
	"fmt"

	"github.com/rbcorrea/meli-challenge/internal/domain/entity"
	"github.com/rbcorrea/meli-challenge/internal/domain/repository"
	"github.com/redis/go-redis/v9"
)

type SearchByCodeUseCase struct {
	redis      *redis.Client
	repository repository.ShortenURLRepository
}

func NewSearchByCodeUseCase(redis *redis.Client, repository repository.ShortenURLRepository) *SearchByCodeUseCase {
	return &SearchByCodeUseCase{
		redis:      redis,
		repository: repository,
	}
}

func (u *SearchByCodeUseCase) Execute(ctx context.Context, code string) (*entity.ShortURL, error) {
	// Buscar no Redis pelo código
	originalURL, err := u.redis.Get(ctx, code).Result()
	if err == nil {
		return &entity.ShortURL{
			OriginalURL: originalURL,
			ShortURL:    "https://me.li/" + code,
			Code:        code,
		}, nil
	}

	// Se não encontrou no Redis, buscar no MongoDB
	url, err := u.repository.FindByCode(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to find URL: %w", err)
	}

	return url, nil
}
