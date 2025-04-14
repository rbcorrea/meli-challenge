package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/rbcorrea/meli-challenge/internal/application/dto"
	"github.com/rbcorrea/meli-challenge/internal/domain/repository"
	"github.com/redis/go-redis/v9"
)

type DeleteURLUseCase struct {
	redis      *redis.Client
	repository repository.ShortenURLRepository
}

func NewDeleteURLUseCase(redis *redis.Client, repository repository.ShortenURLRepository) *DeleteURLUseCase {
	return &DeleteURLUseCase{
		redis:      redis,
		repository: repository,
	}
}

func (u *DeleteURLUseCase) Execute(ctx context.Context, request *dto.DeleteURLRequest) (*dto.DeleteURLResponse, error) {

	url, err := u.repository.FindByCode(ctx, request.Code)
	if err != nil {
		return nil, fmt.Errorf("failed to find URL: %w", err)
	}
	if url == nil {
		return nil, fmt.Errorf("URL not found")
	}

	if err := u.redis.Del(ctx, request.Code).Err(); err != nil {
		return nil, fmt.Errorf("failed to delete from Redis: %w", err)
	}

	update := map[string]interface{}{
		"$set": map[string]interface{}{
			"is_active":  false,
			"deleted_at": time.Now(),
		},
	}
	if err := u.repository.Update(ctx, url.ShortURL, update); err != nil {
		return nil, fmt.Errorf("failed to update MongoDB: %w", err)
	}

	updatedURL, err := u.repository.FindByCode(ctx, request.Code)
	if err != nil {
		return nil, fmt.Errorf("failed to find updated URL: %w", err)
	}

	return dto.NewDeleteURLResponse(
		updatedURL.OriginalURL,
		updatedURL.ShortURL,
		updatedURL.Code,
		updatedURL.CreatedAt,
		updatedURL.IsActive,
		time.Now(),
	), nil
}
