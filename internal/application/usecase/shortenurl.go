package usecase

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/rbcorrea/meli-challenge/internal/application/dto"
	"github.com/rbcorrea/meli-challenge/internal/domain/entity"
	"github.com/rbcorrea/meli-challenge/internal/domain/queue"
	"github.com/rbcorrea/meli-challenge/internal/domain/repository"
	"github.com/redis/go-redis/v9"
)

type ShortenURLUseCase struct {
	repository repository.ShortenURLRepository
	producer   queue.Producer
	redis      *redis.Client
}

func NewShortenURLUseCase(repository repository.ShortenURLRepository, producer queue.Producer, redis *redis.Client) *ShortenURLUseCase {
	return &ShortenURLUseCase{
		repository: repository,
		producer:   producer,
		redis:      redis,
	}
}

func (u *ShortenURLUseCase) Execute(ctx context.Context, request *dto.ShortenURLRequest) (*dto.ShortenURLResponse, error) {
	code := uuid.New().String()
	shortURL := &entity.ShortURL{
		Code:        code,
		OriginalURL: request.URL,
		CreatedAt:   time.Now(),
		IsActive:    true,
		AccessCount: 0,
	}

	response := &dto.ShortenURLResponse{
		OriginalURL: request.URL,
		ShortURL:    "http://localhost:8080/" + code,
		Code:        code,
		CreatedAt:   time.Now(),
		IsActive:    true,
	}

	go func() {
		ctx := context.Background()
		if err := u.producer.PublishShortenURL(ctx, shortURL); err != nil {
			if err := u.repository.Save(ctx, shortURL); err != nil {
				// TODO: Log error
			}
			data, _ := json.Marshal(shortURL)
			if err := u.redis.Set(ctx, code, data, 24*time.Hour).Err(); err != nil {
				// TODO: Log error
			}
		}
	}()

	return response, nil
}
