package models

import (
	"gorm.io/gorm"
	"time"
)

type Url struct {
	gorm.Model
	OriginalURL string
	ShortUrlId  string
	ExpiryDate  time.Time
	UserID      uint
	User        User
}
