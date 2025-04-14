package usecase

import (
	"context"
	"fmt"

	"github.com/rbcorrea/meli-challenge/internal/domain/entity"
	"github.com/rbcorrea/meli-challenge/internal/domain/repository"
)

type SearchByCodeUseCase struct {
	repository repository.ShortenURLRepository
}

func NewSearchByCodeUseCase(repository repository.ShortenURLRepository) *SearchByCodeUseCase {
	return &SearchByCodeUseCase{
		repository: repository,
	}
}

func (u *SearchByCodeUseCase) Execute(ctx context.Context, code string) (*entity.ShortURL, error) {
	url, err := u.repository.FindByCode(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to find URL: %w", err)
	}

	return url, nil
}

func (u *SearchByCodeUseCase) IncrementAccessCount(ctx context.Context, code string) error {
	return u.repository.IncrementAccessCount(ctx, code)
}
