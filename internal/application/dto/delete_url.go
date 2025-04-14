package dto

import "time"

type DeleteURLRequest struct {
	Code string `json:"code" validate:"required"`
}

type DeleteURLResponse struct {
	OriginalURL string    `json:"original_url"`
	ShortURL    string    `json:"short_url"`
	Code        string    `json:"code"`
	CreatedAt   time.Time `json:"created_at"`
	IsActive    bool      `json:"is_active"`
	DeletedAt   time.Time `json:"deleted_at"`
}

func NewDeleteURLResponse(originalURL, shortURL, code string, createdAt time.Time, isActive bool, deletedAt time.Time) *DeleteURLResponse {
	return &DeleteURLResponse{
		OriginalURL: originalURL,
		ShortURL:    shortURL,
		Code:        code,
		CreatedAt:   createdAt,
		IsActive:    isActive,
		DeletedAt:   deletedAt,
	}
}
