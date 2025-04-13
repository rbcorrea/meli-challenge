package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/rbcorrea/meli-challenge/internal/domain/entity"
)

type ShortenURLUseCase struct {
}

func (u *ShortenURLUseCase) Execute(ctx context.Context, originalURL string) (*entity.ShortURL, error) {
	code := uuid.New().String()
	shortURL := entity.NewShortURL(originalURL, code)

	// err := u.Producer.PublishShortenURL(ctx, shortURL)
	// if err != nil {
	// 	return nil, err
	// }

	return shortURL, nil
}
