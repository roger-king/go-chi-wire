package config

import (
	"os"
)

// Config - top level config for server
type Config struct {
	Server *ServerConfig
}

// ServerConfig - configration of the server
type ServerConfig struct {
	GoENV string
	Port  string
}

// GetConfig - gets configuration
func GetConfig() *Config {
	return &Config{
		Server: &ServerConfig{
			GoENV: EnvOrDefaultString("GO_ENV", "local"),
			Port:  EnvOrDefaultString("PORT", "5000"),
		},
	}
}

// EnvOrDefaultString - Returns the environment variable or a prefixed default value (e.g. port)
func EnvOrDefaultString(envVar string, defaultValue string) string {
	value := os.Getenv(envVar)
	if value == "" {
		return defaultValue
	}

	return value
}
