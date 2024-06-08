package config

type SecurityConfig struct {
	JWT struct {
		SecretKey string `yaml:"secret_key"`
		Issuer    string `yaml:"issuer"`
	}
}
