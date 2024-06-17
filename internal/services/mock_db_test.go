package services

import (
	appconfig "github.com/EmmanuelStan12/URLShortner/internal/config"
	"github.com/EmmanuelStan12/URLShortner/internal/database"
	"gorm.io/gorm"
	"testing"
)

func InitTestDB(t *testing.T) *gorm.DB {
	conf, err := appconfig.InitConfig("../../app_config_dev.yml")
	if err != nil {
		t.Errorf("Test failed with error, %s.", err)
	}

	db, err := database.InitDatabase(conf.DB)
	if err != nil {
		t.Errorf("Test failed with error, %s", err)
	}

	return db
}

func TeardownTestDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("DROP TABLE IF EXISTS users")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}
