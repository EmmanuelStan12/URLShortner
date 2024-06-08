package database

import (
	"github.com/EmmanuelStan12/URLShortner/internal/config"
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

	return db, nil
}
