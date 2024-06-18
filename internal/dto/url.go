package dto

import (
	"time"
)

type UrlDTO struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ExpiryDate  time.Time `json:"expiry_date"`
	OriginalURL string    `json:"original_url"`
	ShortURL    string    `json:"short_url"`
	UserID      uint      `json:"user_id"`
}

type CreateShortUrl struct {
	OriginalURL string `json:"original_url"`
}

type UrlClickDTO struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserAgent string    `json:"user_agent"`
	IPAddress string    `json:"ip_address"`
	UrlID     uint      `json:"url_id"`
}
