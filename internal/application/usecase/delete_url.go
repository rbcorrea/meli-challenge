package usecase

import (
	"context"
	"fmt"

	"github.com/rbcorrea/meli-challenge/internal/application/dto"
	"github.com/rbcorrea/meli-challenge/internal/domain/entity"
	"github.com/rbcorrea/meli-challenge/internal/domain/repository"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
)

type DeleteURLUseCase struct {
	repository repository.ShortenURLRepository
	redis      redis.Cmdable
}

func NewDeleteURLUseCase(repository repository.ShortenURLRepository, redis redis.Cmdable) *DeleteURLUseCase {
	return &DeleteURLUseCase{
		repository: repository,
		redis:      redis,
	}
}

func (u *DeleteURLUseCase) Execute(ctx context.Context, request *dto.DeleteURLRequest) (*entity.ShortURL, error) {
	url, err := u.repository.FindByCode(ctx, request.Code)
	if err != nil {
		return nil, fmt.Errorf("failed to find URL: %w", err)
	}

	if url == nil {
		return nil, fmt.Errorf("URL not found")
	}

	update := bson.M{
		"$set": bson.M{
			"is_active": false,
		},
	}
	if err := u.repository.Update(ctx, request.Code, update); err != nil {
		return nil, fmt.Errorf("failed to update MongoDB: %w", err)
	}

	if err := u.redis.Del(ctx, request.Code).Err(); err != nil {
		return nil, fmt.Errorf("failed to delete from Redis: %w", err)
	}

	url.IsActive = false

	return url, nil
}
