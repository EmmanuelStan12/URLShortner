package config

import (
	"fmt"
	"github.com/EmmanuelStan12/URLShortner/internal/util"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	DB       DBConfig       `yaml:"db"`
	Security SecurityConfig `yaml:"security"`
}

func InitRootConfig() (*Config, error) {
	env := os.Getenv(util.ENVIRONMENT)
	path := fmt.Sprintf("app_config_%s.yml", env)
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	config, err := InitConfig(absPath)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func InitConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	decoder := yaml.NewDecoder(file)

	config := &Config{}
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}
	return config, nil
}
