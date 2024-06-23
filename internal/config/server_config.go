package config

type ServerConfig struct {
	Port     string `yaml:"port"`
	Hostname string `yaml:"hostname"`
	Profile  string `yaml:"profile"`
}
