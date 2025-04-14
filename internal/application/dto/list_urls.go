package dto

type ListURLsRequest struct {
	Page     int `json:"page" query:"page" validate:"min=1"`
	PageSize int `json:"page_size" query:"page_size" validate:"min=1,max=100"`
}

type URLMapping struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type ListURLsResponse struct {
	URLs       []URLMapping `json:"urls"`
	Total      int          `json:"total"`
	Page       int          `json:"page"`
	PageSize   int          `json:"page_size"`
	TotalPages int          `json:"total_pages"`
}
