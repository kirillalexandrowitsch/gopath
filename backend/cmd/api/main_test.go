package main

import (
	"context"
	"testing"

	"github.com/kirillalexandrowitsch/gopath/backend/internal/database"
	"github.com/kirillalexandrowitsch/gopath/backend/internal/learning"
)

func TestLearningStoreFromEnvUsesMemoryStoreByDefault(t *testing.T) {
	t.Setenv(database.EnvDatabaseURL, "")

	store, closeStore, err := learningStoreFromEnv(context.Background())
	if err != nil {
		t.Fatalf("create learning store: %v", err)
	}
	if err := closeStore(); err != nil {
		t.Fatalf("close learning store: %v", err)
	}

	if _, ok := store.(*learning.MemoryStore); !ok {
		t.Fatalf("expected *learning.MemoryStore, got %T", store)
	}
}
