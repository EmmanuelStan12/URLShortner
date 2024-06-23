package database

import (
	"fmt"
	"github.com/EmmanuelStan12/URLShortner/internal/config"
	"github.com/EmmanuelStan12/URLShortner/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase(dbConfig config.DBConfig) (*gorm.DB, error) {
	dsn := dbConfig.ConstructDSN()
	db, err := gorm.Open(
		mysql.New(
			mysql.Config{
				DSN: dsn,
			},
		),
		&gorm.Config{},
	)
	fmt.Printf("Logging into db with dsn: %s\n", dsn)

	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&models.User{}, &models.Url{}, &models.UrlClick{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
