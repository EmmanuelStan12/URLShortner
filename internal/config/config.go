package config

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	DB       DBConfig       `yaml:"db"`
	Security SecurityConfig `yaml:"security"`
}

func InitRootConfig() (*Config, error) {
	absPath, err := filepath.Abs("app_config.yml")
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
	data, err := loadConfigPreprocessor(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	if err := yaml.Unmarshal([]byte(data), &config); err != nil {
		return nil, err
	}
	return config, nil
}

func loadConfigPreprocessor(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	content := string(data)
	envExp := regexp.MustCompile(`\$\{(\w+)\}`)

	processedContent := envExp.ReplaceAllStringFunc(content, func(match string) string {
		envVar := envExp.FindStringSubmatch(match)[1]
		return os.Getenv(envVar)
	})

	return processedContent, nil
}
