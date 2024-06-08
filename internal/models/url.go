package models

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	OriginalURL string
	ShortURL    string
	UserID      uint
	User        User
}
