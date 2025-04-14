package mongo

import (
	"context"
	"fmt"

	"github.com/rbcorrea/meli-challenge/internal/domain/entity"
	"github.com/rbcorrea/meli-challenge/internal/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoShortenURLRepository struct {
	collection *mongo.Collection
}

func NewMongoShortenURLRepository(collection *mongo.Collection) repository.ShortenURLRepository {
	return &MongoShortenURLRepository{
		collection: collection,
	}
}

func (r *MongoShortenURLRepository) FindByOriginalURL(ctx context.Context, originalURL string) (*entity.ShortURL, error) {
	var shortURL entity.ShortURL
	err := r.collection.FindOne(ctx, bson.M{
		"original_url": originalURL,
		"is_active":    true,
	}).Decode(&shortURL)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &shortURL, nil
}

func (r *MongoShortenURLRepository) Save(ctx context.Context, shortURL *entity.ShortURL) error {
	filter := bson.M{"short_url": shortURL.ShortURL}
	update := bson.M{
		"$set": bson.M{
			"original_url": shortURL.OriginalURL,
			"code":         shortURL.Code,
			"created_at":   shortURL.CreatedAt,
			"is_active":    shortURL.IsActive,
		},
	}
	opts := options.Update().SetUpsert(true)
	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	return err
}

func (r *MongoShortenURLRepository) FindByShortURL(ctx context.Context, shortURL string) (*entity.ShortURL, error) {
	filter := bson.M{
		"short_url": shortURL,
		"is_active": true,
	}
	var result entity.ShortURL
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *MongoShortenURLRepository) FindByCode(ctx context.Context, code string) (*entity.ShortURL, error) {
	filter := bson.M{
		"code":      code,
		"is_active": true,
	}
	var result entity.ShortURL
	err := r.collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *MongoShortenURLRepository) Update(ctx context.Context, shortURL string, update interface{}) error {
	filter := bson.M{"short_url": shortURL}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *MongoShortenURLRepository) IncrementAccessCount(ctx context.Context, code string) error {
	filter := bson.M{"code": code}
	update := bson.M{
		"$inc": bson.M{"access_count": 1},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to increment access count: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("url not found")
	}

	return nil
}
