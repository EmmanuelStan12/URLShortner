package services

import (
	appconfig "github.com/EmmanuelStan12/URLShortner/internal/config"
	"github.com/EmmanuelStan12/URLShortner/internal/database"
	"github.com/EmmanuelStan12/URLShortner/internal/models"
	"gorm.io/gorm"
	"testing"
)

func initTestDB(t *testing.T) (*gorm.DB, appconfig.Config) {
	conf, err := appconfig.InitConfig("../../app_config_dev.yml")
	if err != nil {
		t.Errorf("Test failed with error, %s.", err)
	}

	db, err := database.InitDatabase(conf.DB)
	if err != nil {
		t.Errorf("Test failed with error, %s", err)
	}
	initMigrations(db)

	return db, *conf
}

func initMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Url{}, &models.UrlClick{})
	if err != nil {
		return
	}
}

func teardownTestDB(db *gorm.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("DROP TABLE IF EXISTS users")
	db.Exec("DROP TABLE IF EXISTS urls")
	db.Exec("DROP TABLE IF EXISTS url_clicks")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func handlePanic[T any](t *testing.T, fn func()) {
	defer func() {
		if r := recover(); r != nil {
			_, ok := r.(T)
			if !ok {
				t.Errorf("expected an error, got %v", r)
			}
			t.Logf("error: %v", r)
		} else {
			t.Error("expected panic but did not panic")
		}
	}()

	fn()
}
