package database

import (
	"context"
	"errors"
	"testing"
)

func TestConfigFromEnv(t *testing.T) {
	databaseURL := "postgres://gopath:gopath_dev_password@localhost:5432/gopath?sslmode=disable"
	t.Setenv(EnvDatabaseURL, databaseURL)

	config := ConfigFromEnv()

	if config.URL != databaseURL {
		t.Fatalf("expected database URL %q, got %q", databaseURL, config.URL)
	}
}

func TestConfigValidateMissingURL(t *testing.T) {
	t.Parallel()

	config := Config{}

	if err := config.Validate(); !errors.Is(err, ErrMissingDatabaseURL) {
		t.Fatalf("expected ErrMissingDatabaseURL, got %v", err)
	}
}

func TestOpenRequiresDatabaseURL(t *testing.T) {
	t.Parallel()

	db, err := Open(context.Background(), Config{})
	if db != nil {
		t.Fatal("expected nil database handle")
	}
	if !errors.Is(err, ErrMissingDatabaseURL) {
		t.Fatalf("expected ErrMissingDatabaseURL, got %v", err)
	}
}
