// Package config
package config

import (
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	Env     string
	Version string

	ReferenceDate int
}

// LoadConfig loads the configuration from the environment variables.
// if the .env file is not found, it will use the default values.
func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		slog.Warn("Error loading .env", "error", err)
	}

	currYear := time.Now().Year()

	cfg := &Config{
		Port:    os.Getenv("PORT"),
		Env:     os.Getenv("ENV"),
		Version: os.Getenv("VERSION"),

		ReferenceDate: currYear,
	}

	err = cfg.validate()
	if err != nil {
		slog.Warn("Validation error", "error", err)
	}

	return cfg, nil
}

// validate validates the configuration.
// returning default values if the environment variables are not set.
// returning an error if version is not set.
func (c *Config) validate() error {
	const (
		PORT    = "3000"
		ENV     = "dev"
		VERSION = "debug"
	)

	if c.Port == "" {
		slog.Info("Missing PORT environment variable using default", "PORT", PORT)
		c.Port = PORT
	}

	if c.Env == "" {
		slog.Info("Missing ENV environment variable using default", "ENV", ENV)
		c.Env = ENV
	}

	if c.Version == "" || c.Env != "prod" {
		slog.Info(
			"Missing VERSION environment variable using default",
			"VERSION",
			VERSION,
		)
		c.Version = VERSION
	}

	return nil
}
