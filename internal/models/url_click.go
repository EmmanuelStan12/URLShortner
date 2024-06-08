package models

import "gorm.io/gorm"

type UrlClick struct {
	gorm.Model
	UserAgent string
	IPAddress string
	UrlID     uint
	Url       Url
}
