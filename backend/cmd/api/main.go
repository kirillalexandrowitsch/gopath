package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kirillalexandrowitsch/gopath/backend/internal/httpserver"
)

func main() {
	addr := env("GOPATH_API_ADDR", ":8080")

	server := &http.Server{
		Addr:              addr,
		Handler:           httpserver.NewRouter(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("starting GoPath API on %s", addr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server stopped: %v", err)
	}
}

func env(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
