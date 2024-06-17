package database

import (
	"github.com/EmmanuelStan12/URLShortner/internal/config"
	"testing"
)

func TestInitDatabase(t *testing.T) {
	conf, err := config.InitConfig("../../app_config_test.yml")
	if err != nil {
		t.Errorf("Failed to load yml config, %s.", err.Error())
	}

	_, err = InitDatabase(conf.DB)
	if err != nil {
		t.Errorf("Failed to connect to db, %s.", err.Error())
	}
}
