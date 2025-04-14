package entity

import "time"

type ShortURL struct {
	OriginalURL string    `json:"original_url" bson:"original_url"`
	ShortURL    string    `json:"short_url" bson:"short_url"`
	Code        string    `json:"code" bson:"code"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	IsActive    bool      `json:"is_active" bson:"is_active"`
}

func NewShortURL(originalURL, code string) *ShortURL {
	return &ShortURL{
		OriginalURL: originalURL,
		Code:        code,
		ShortURL:    "https://me.li/" + code,
		CreatedAt:   time.Now(),
		IsActive:    true,
	}
}
