package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rbcorrea/meli-challenge/internal/application/dto"
	"github.com/rbcorrea/meli-challenge/internal/domain/entity"
	"github.com/rbcorrea/meli-challenge/internal/domain/queue"
	"github.com/rbcorrea/meli-challenge/internal/domain/repository"
)

type ShortenURLUseCase struct {
	repository repository.ShortenURLRepository
	producer   queue.Producer
}

func NewShortenURLUseCase(repository repository.ShortenURLRepository, producer queue.Producer) *ShortenURLUseCase {
	return &ShortenURLUseCase{
		repository: repository,
		producer:   producer,
	}
}

func (u *ShortenURLUseCase) Execute(ctx context.Context, request *dto.ShortenURLRequest) (*dto.ShortenURLResponse, error) {
	existingURL, err := u.repository.FindByOriginalURL(ctx, request.URL)
	if err != nil {
		return nil, err
	}

	if existingURL != nil {
		return dto.NewShortenURLResponse(existingURL.OriginalURL, existingURL.ShortURL, existingURL.Code, time.Now(), true), nil
	}

	code := uuid.New().String()
	shortURL := entity.NewShortURL(request.URL, code)

	err = u.repository.Save(ctx, shortURL)
	if err != nil {
		return nil, err
	}

	err = u.producer.PublishShortenURL(ctx, shortURL)
	if err != nil {
		// TODO: Implementar log
	}

	return dto.NewShortenURLResponse(shortURL.OriginalURL, shortURL.ShortURL, shortURL.Code, shortURL.CreatedAt, shortURL.IsActive), nil
}
