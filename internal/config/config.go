// Package config
package config

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

const (
	defaultPORT       = "3000"
	defaultMaxResults = 1000
)

type Config struct {
	Port    string
	Env     string
	Version string

	MaxResults int
}

// LoadConfig loads the configuration from the environment variables.
// if the .env file is not found, it will use the default values.
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil && !errors.Is(err, fs.ErrNotExist) {
		slog.Warn("Error loading .env", "error", err)
	}

	cfg := &Config{
		Port:    os.Getenv("PORT"),
		Env:     os.Getenv("ENV"),
		Version: os.Getenv("VERSION"),

		MaxResults: defaultMaxResults,
	}

	err := cfg.validate()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// validate validates the configuration.
// ENV must be "dev" or "prod"; VERSION is required in prod.
// PORT defaults to 3000 if unset.
func (c *Config) validate() error {
	switch c.Env {
	case "":
		return fmt.Errorf("missing 'ENV' environment variable")
	case "dev":
		// INFO: When in "dev" env, we force VERSION to "debug"
		// override the literal below if testing `Info.Version`
		c.Version = "debug"
		slog.Info(
			"Detected environment: 'DEV', running in debug mode",
			"VERSION",
			c.Version,
		)
	case "prod":
		if c.Version == "" {
			return fmt.Errorf("missing 'VERSION' environment variable")
		}
	default:
		return fmt.Errorf("invalid 'ENV' environment variable: %q", c.Env)
	}

	if c.Port == "" {
		slog.Info(
			"Missing PORT environment variable using default",
			"PORT",
			defaultPORT,
		)
		c.Port = defaultPORT
	}

	return nil
}
