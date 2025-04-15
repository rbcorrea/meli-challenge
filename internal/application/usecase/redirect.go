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

type RedirectUseCase struct {
	redis      *redis.Client
	repository repository.ShortenURLRepository
}

func NewRedirectUseCase(repository repository.ShortenURLRepository, redis *redis.Client) *RedirectUseCase {
	return &RedirectUseCase{
		redis:      redis,
		repository: repository,
	}
}

func (u *RedirectUseCase) Execute(ctx context.Context, code string) (*entity.ShortURL, error) {
	redisData, err := u.redis.Get(ctx, code).Result()
	if err == nil {
		var shortURL entity.ShortURL
		if err := json.Unmarshal([]byte(redisData), &shortURL); err == nil {
			if !shortURL.IsActive {
				return nil, fmt.Errorf("URL not found")
			}
			return &shortURL, nil
		}
	}

	url, err := u.repository.FindByCode(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to find URL: %w", err)
	}
	if url == nil || !url.IsActive {
		return nil, fmt.Errorf("URL not found")
	}

	return url, nil
}

func (u *RedirectUseCase) IncrementAccessCount(ctx context.Context, code string) error {
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
