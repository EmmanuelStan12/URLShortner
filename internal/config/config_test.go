package config

import "testing"

func TestInitConfig(t *testing.T) {
	t.Run("Valid file", func(t *testing.T) {
		config, err := InitConfig("../../app_config_test.yml")
		if err != nil {
			t.Errorf("Test failed with error, %s.", err)
		}

		t.Logf("Test passed with config %v.", config)
	})

	t.Run("Invalid file", func(t *testing.T) {
		config, err := InitConfig("../../app_config_incorrect.yml")
		if err == nil {
			t.Errorf("Expected error but got, %v.", config)
		}
	})
}

func TestInitRootConfig(t *testing.T) {
	config, err := InitRootConfig()
	if err != nil {
		t.Errorf("Test failed with error, %s.", err)
	}

	t.Logf("Test passed with config %v.", config)
}
