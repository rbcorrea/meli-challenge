package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rbcorrea/meli-challenge/internal/domain/entity"
	"github.com/rbcorrea/meli-challenge/internal/domain/repository"
	"github.com/redis/go-redis/v9"
)

type SearchByCodeUseCase struct {
	repository repository.ShortenURLRepository
	redis      *redis.Client
}

func NewSearchByCodeUseCase(repository repository.ShortenURLRepository, redis *redis.Client) *SearchByCodeUseCase {
	return &SearchByCodeUseCase{
		repository: repository,
		redis:      redis,
	}
}

func (u *SearchByCodeUseCase) Execute(ctx context.Context, code string) (*entity.ShortURL, error) {
	redisData, err := u.redis.Get(ctx, code).Result()
	if err == nil {
		var shortURL entity.ShortURL
		if err := json.Unmarshal([]byte(redisData), &shortURL); err == nil {
			return &shortURL, nil
		}
	}

	shortURL, err := u.repository.FindByCode(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to find URL: %w", err)
	}

	if shortURL != nil && shortURL.IsActive {
		go func() {
			ctx := context.Background()
			data, _ := json.Marshal(shortURL)
			u.redis.Set(ctx, code, data, 24*time.Hour)
		}()
	}

	return shortURL, nil
}

func (u *SearchByCodeUseCase) IncrementAccessCount(ctx context.Context, code string) error {
	redisData, err := u.redis.Get(ctx, code).Result()
	if err == nil {
		var shortURL entity.ShortURL
		if err := json.Unmarshal([]byte(redisData), &shortURL); err == nil {
			shortURL.AccessCount++
			data, _ := json.Marshal(shortURL)
			u.redis.Set(ctx, code, data, 24*time.Hour)
		}
	}
	return u.repository.IncrementAccessCount(ctx, code)
}
