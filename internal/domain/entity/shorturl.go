package dto

type ShortURL struct {
	OriginalURL string
	ShortURL    string
}

func NewShortURL(originalURL, shortURL string) *ShortURL {
	return &ShortURL{
		OriginalURL: originalURL,
		ShortURL:    shortURL,
	}
}
