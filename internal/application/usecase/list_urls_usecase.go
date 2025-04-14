package usecase

import (
	"context"
	"fmt"

	"github.com/rbcorrea/meli-challenge/internal/application/dto"
	"github.com/redis/go-redis/v9"
)

type ListURLsUseCase struct {
	redis *redis.Client
}

func NewListURLsUseCase(redis *redis.Client) *ListURLsUseCase {
	return &ListURLsUseCase{
		redis: redis,
	}
}

func (u *ListURLsUseCase) Execute(ctx context.Context, request *dto.ListURLsRequest) (*dto.ListURLsResponse, error) {
	offset := (request.Page - 1) * request.PageSize

	keys, err := u.redis.Keys(ctx, "*").Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get keys: %w", err)
	}

	total := len(keys)
	totalPages := (total + request.PageSize - 1) / request.PageSize

	if request.Page > totalPages {
		request.Page = totalPages
		offset = (request.Page - 1) * request.PageSize
	}

	start := offset
	end := offset + request.PageSize
	if end > total {
		end = total
	}

	pageKeys := keys[start:end]

	urls := make([]dto.URLMapping, 0, len(pageKeys))
	for _, key := range pageKeys {
		value, err := u.redis.Get(ctx, key).Result()
		if err != nil {
			return nil, fmt.Errorf("failed to get value for key %s: %w", key, err)
		}

		urls = append(urls, dto.URLMapping{
			ShortURL:    key,
			OriginalURL: value,
		})
	}

	return &dto.ListURLsResponse{
		URLs:       urls,
		Total:      total,
		Page:       request.Page,
		PageSize:   request.PageSize,
		TotalPages: totalPages,
	}, nil
}
