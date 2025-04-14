package queue

import (
	"context"

	"github.com/rbcorrea/meli-challenge/internal/domain/entity"
)

type Producer interface {
	PublishShortenURL(ctx context.Context, shortURL *entity.ShortURL) error
}
