//go:build integration

package sessions

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/kirillalexandrowitsch/gopath/backend/internal/database/dbtest"
	"github.com/kirillalexandrowitsch/gopath/backend/internal/users"
)

func TestRepositoryCreate(t *testing.T) {
	repositories := setupRepositories(t)
	ctx := context.Background()
	user := createUserFixture(t, repositories.users)
	expiresAt := time.Now().UTC().Add(time.Hour).Truncate(time.Microsecond)
	lastUsedAt := time.Now().UTC().Truncate(time.Microsecond)

	session, err := repositories.sessions.Create(ctx, CreateParams{
		ID:         "test-session-create",
		UserID:     user.ID,
		TokenHash:  "test-token-hash-create",
		ExpiresAt:  expiresAt,
		LastUsedAt: &lastUsedAt,
	})
	if err != nil {
		t.Fatalf("create session: %v", err)
	}

	if session.ID != "test-session-create" {
		t.Fatalf("expected id %q, got %q", "test-session-create", session.ID)
	}
	if session.UserID != user.ID {
		t.Fatalf("expected user id %q, got %q", user.ID, session.UserID)
	}
	if session.TokenHash != "test-token-hash-create" {
		t.Fatalf("expected token hash %q, got %q", "test-token-hash-create", session.TokenHash)
	}
	if !session.ExpiresAt.Equal(expiresAt) {
		t.Fatalf("expected expires at %s, got %s", expiresAt, session.ExpiresAt)
	}
	if session.LastUsedAt == nil {
		t.Fatal("expected last used timestamp")
	}
	if !session.LastUsedAt.Equal(lastUsedAt) {
		t.Fatalf("expected last used at %s, got %s", lastUsedAt, *session.LastUsedAt)
	}
	if session.CreatedAt.IsZero() {
		t.Fatal("expected created timestamp")
	}
	if session.UpdatedAt.IsZero() {
		t.Fatal("expected updated timestamp")
	}
}

func TestRepositoryFindByTokenHash(t *testing.T) {
	repositories := setupRepositories(t)
	ctx := context.Background()
	user := createUserFixture(t, repositories.users)

	createdSession := createSessionFixture(t, repositories.sessions, CreateParams{
		ID:        "test-session-find-token",
		UserID:    user.ID,
		TokenHash: "test-token-hash-find",
		ExpiresAt: time.Now().UTC().Add(time.Hour).Truncate(time.Microsecond),
	})

	foundSession, err := repositories.sessions.FindByTokenHash(ctx, createdSession.TokenHash)
	if err != nil {
		t.Fatalf("find session by token hash: %v", err)
	}

	assertSameSession(t, createdSession, foundSession)
}

func TestRepositoryDelete(t *testing.T) {
	repositories := setupRepositories(t)
	ctx := context.Background()
	user := createUserFixture(t, repositories.users)

	createdSession := createSessionFixture(t, repositories.sessions, CreateParams{
		ID:        "test-session-delete",
		UserID:    user.ID,
		TokenHash: "test-token-hash-delete",
		ExpiresAt: time.Now().UTC().Add(time.Hour).Truncate(time.Microsecond),
	})

	if err := repositories.sessions.Delete(ctx, createdSession.ID); err != nil {
		t.Fatalf("delete session: %v", err)
	}

	if _, err := repositories.sessions.FindByTokenHash(ctx, createdSession.TokenHash); !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound after delete, got %v", err)
	}
}

func TestRepositoryMissingSession(t *testing.T) {
	repositories := setupRepositories(t)
	ctx := context.Background()

	if _, err := repositories.sessions.FindByTokenHash(ctx, "missing-token-hash"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound for missing token hash, got %v", err)
	}

	if err := repositories.sessions.Delete(ctx, "missing-session"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound for missing delete target, got %v", err)
	}
}

type testRepositories struct {
	sessions *Repository
	users    *users.Repository
}

func setupRepositories(t *testing.T) testRepositories {
	t.Helper()

	db := dbtest.Open(t)
	applyMigrationsIfNeeded(t, db)
	dbtest.TruncateTables(t, db, "sessions", "users")

	return testRepositories{
		sessions: NewRepository(db),
		users:    users.NewRepository(db),
	}
}

func applyMigrationsIfNeeded(t testing.TB, db *sql.DB) {
	t.Helper()

	if tableExists(t, db, "sessions") {
		return
	}
	if !tableExists(t, db, "users") {
		dbtest.ApplyMigrations(t, db)
		return
	}

	applyMigrationFile(t, db, "000004_sessions_schema.sql")
}

func tableExists(t testing.TB, db *sql.DB, tableName string) bool {
	t.Helper()

	var exists bool
	err := db.QueryRow(`
		SELECT EXISTS (
			SELECT 1
			FROM information_schema.tables
			WHERE table_schema = 'public'
				AND table_name = $1
		)
	`, tableName).Scan(&exists)
	if err != nil {
		t.Fatalf("check table %q: %v", tableName, err)
	}

	return exists
}

func applyMigrationFile(t testing.TB, db *sql.DB, name string) {
	t.Helper()

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("resolve sessions test path")
	}

	migrationPath := filepath.Join(filepath.Dir(filename), "..", "..", "migrations", name)
	sqlBytes, err := os.ReadFile(migrationPath)
	if err != nil {
		t.Fatalf("read migration %s: %v", name, err)
	}

	if _, err := db.Exec(string(sqlBytes)); err != nil {
		t.Fatalf("apply migration %s: %v", name, err)
	}
}

func createUserFixture(t *testing.T, repository *users.Repository) users.User {
	t.Helper()

	user, err := repository.Create(context.Background(), users.CreateParams{
		ID:           "test-user-session",
		Email:        "session-user@example.com",
		Handle:       "@sessionuser",
		DisplayName:  "Session User",
		PasswordHash: "hashed-password",
	})
	if err != nil {
		t.Fatalf("create user fixture: %v", err)
	}

	return user
}

func createSessionFixture(t *testing.T, repository *Repository, params CreateParams) Session {
	t.Helper()

	session, err := repository.Create(context.Background(), params)
	if err != nil {
		t.Fatalf("create session fixture: %v", err)
	}

	return session
}

func assertSameSession(t *testing.T, expected Session, actual Session) {
	t.Helper()

	if actual.ID != expected.ID {
		t.Fatalf("expected id %q, got %q", expected.ID, actual.ID)
	}
	if actual.UserID != expected.UserID {
		t.Fatalf("expected user id %q, got %q", expected.UserID, actual.UserID)
	}
	if actual.TokenHash != expected.TokenHash {
		t.Fatalf("expected token hash %q, got %q", expected.TokenHash, actual.TokenHash)
	}
	if !actual.ExpiresAt.Equal(expected.ExpiresAt) {
		t.Fatalf("expected expires at %s, got %s", expected.ExpiresAt, actual.ExpiresAt)
	}
	if (actual.LastUsedAt == nil) != (expected.LastUsedAt == nil) {
		t.Fatalf("expected last used at %v, got %v", expected.LastUsedAt, actual.LastUsedAt)
	}
	if actual.LastUsedAt != nil && !actual.LastUsedAt.Equal(*expected.LastUsedAt) {
		t.Fatalf("expected last used at %s, got %s", *expected.LastUsedAt, *actual.LastUsedAt)
	}
}
