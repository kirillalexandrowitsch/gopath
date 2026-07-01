package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kirillalexandrowitsch/gopath/backend/internal/database"
	"github.com/kirillalexandrowitsch/gopath/backend/internal/httpserver"
	"github.com/kirillalexandrowitsch/gopath/backend/internal/learning"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatalf("api stopped: %v", err)
	}
}

func run(ctx context.Context) error {
	addr := env("GOPATH_API_ADDR", ":8080")
	learningStore, closeStore, err := learningStoreFromEnv(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err := closeStore(); err != nil {
			log.Printf("close learning store: %v", err)
		}
	}()

	server := &http.Server{
		Addr:              addr,
		Handler:           httpserver.NewRouterWithStore(learningStore),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("starting GoPath API on %s", addr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("server stopped: %w", err)
	}

	return nil
}

func learningStoreFromEnv(ctx context.Context) (learning.Store, func() error, error) {
	config := database.ConfigFromEnv()
	if config.URL == "" {
		return learning.NewMemoryStore(), func() error { return nil }, nil
	}

	db, err := database.Open(ctx, config)
	if err != nil {
		return nil, nil, fmt.Errorf("connect PostgreSQL learning store: %w", err)
	}

	return learning.NewPostgresStore(db), db.Close, nil
}

func env(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
