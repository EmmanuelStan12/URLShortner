package database

import (
	"github.com/EmmanuelStan12/URLShortner/internal/config"
	"github.com/EmmanuelStan12/URLShortner/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase(dbConfig config.DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open(
		mysql.New(
			mysql.Config{
				DSN: dbConfig.ConstructDSN(),
			},
		),
		&gorm.Config{},
	)

	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.User{}, &models.Url{}, &models.UrlClick{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
