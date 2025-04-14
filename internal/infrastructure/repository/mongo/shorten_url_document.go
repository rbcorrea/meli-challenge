package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShortURLDocument struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	OriginalURL string             `bson:"original_url"`
	ShortURL    string             `bson:"short_url"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}
