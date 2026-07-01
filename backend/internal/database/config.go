package database

import (
	"errors"
	"os"
)

const EnvDatabaseURL = "GOPATH_DATABASE_URL"

var ErrMissingDatabaseURL = errors.New("GOPATH_DATABASE_URL is required")

type Config struct {
	URL string
}

func ConfigFromEnv() Config {
	return Config{
		URL: os.Getenv(EnvDatabaseURL),
	}
}

func (c Config) Validate() error {
	if c.URL == "" {
		return ErrMissingDatabaseURL
	}

	return nil
}
