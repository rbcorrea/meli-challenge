package dto

import "time"

type ShortenURLRequest struct {
	URL string `json:"url" validate:"required,url"`
}

type ShortenURLResponse struct {
	OriginalURL string    `json:"original_url"`
	ShortURL    string    `json:"short_url"`
	Code        string    `json:"code"`
	CreatedAt   time.Time `json:"created_at"`
	IsActive    bool      `json:"is_active"`
}

func NewShortenURLResponse(originalURL, shortURL, code string, createdAt time.Time, isActive bool) *ShortenURLResponse {
	return &ShortenURLResponse{
		OriginalURL: originalURL,
		ShortURL:    shortURL,
		Code:        code,
		CreatedAt:   createdAt,
		IsActive:    isActive,
	}
}
