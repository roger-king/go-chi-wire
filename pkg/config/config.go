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

// NewConfig - Provider for returning application config.
func NewConfig() *Config {
	return &Config{
		Server: &ServerConfig{
			GoENV: envOrDefaultString("GO_ENV", "local"),
			Port:  envOrDefaultString("PORT", "5000"),
		},
	}
}

// EnvOrDefaultString - Returns the environment variable or a prefixed default value (e.g. port)
func envOrDefaultString(envVar string, defaultValue string) string {
	value := os.Getenv(envVar)
	if value == "" {
		return defaultValue
	}

	return value
}
