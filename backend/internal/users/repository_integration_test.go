//go:build integration

package users

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/kirillalexandrowitsch/gopath/backend/internal/database/dbtest"
)

func TestRepositoryCreate(t *testing.T) {
	repository := setupRepository(t)
	ctx := context.Background()

	user, err := repository.Create(ctx, CreateParams{
		ID:           "test-user-create",
		Email:        "create@example.com",
		Handle:       "@create",
		DisplayName:  "Create User",
		PasswordHash: "hashed-password",
	})
	if err != nil {
		t.Fatalf("create user: %v", err)
	}

	if user.ID != "test-user-create" {
		t.Fatalf("expected id %q, got %q", "test-user-create", user.ID)
	}
	if user.Email != "create@example.com" {
		t.Fatalf("expected email %q, got %q", "create@example.com", user.Email)
	}
	if user.Handle != "@create" {
		t.Fatalf("expected handle %q, got %q", "@create", user.Handle)
	}
	if user.DisplayName != "Create User" {
		t.Fatalf("expected display name %q, got %q", "Create User", user.DisplayName)
	}
	if user.PasswordHash != "hashed-password" {
		t.Fatalf("expected password hash %q, got %q", "hashed-password", user.PasswordHash)
	}
	if user.Role != RoleLearner {
		t.Fatalf("expected default role %q, got %q", RoleLearner, user.Role)
	}
	if user.CreatedAt.IsZero() {
		t.Fatal("expected created timestamp")
	}
	if user.UpdatedAt.IsZero() {
		t.Fatal("expected updated timestamp")
	}
}

func TestRepositoryFindByID(t *testing.T) {
	repository := setupRepository(t)
	ctx := context.Background()

	createdUser := createUser(t, repository, CreateParams{
		ID:           "test-user-find-id",
		Email:        "find-id@example.com",
		Handle:       "@findid",
		DisplayName:  "Find ID User",
		PasswordHash: "hashed-password",
		Role:         RoleAdmin,
	})

	foundUser, err := repository.FindByID(ctx, createdUser.ID)
	if err != nil {
		t.Fatalf("find user by id: %v", err)
	}

	assertSameUser(t, createdUser, foundUser)
}

func TestRepositoryFindByEmail(t *testing.T) {
	repository := setupRepository(t)
	ctx := context.Background()

	createdUser := createUser(t, repository, CreateParams{
		ID:           "test-user-find-email",
		Email:        "find-email@example.com",
		Handle:       "@findemail",
		DisplayName:  "Find Email User",
		PasswordHash: "hashed-password",
		Role:         RoleLearner,
	})

	foundUser, err := repository.FindByEmail(ctx, createdUser.Email)
	if err != nil {
		t.Fatalf("find user by email: %v", err)
	}

	assertSameUser(t, createdUser, foundUser)
}

func TestRepositoryFindMissingUser(t *testing.T) {
	repository := setupRepository(t)
	ctx := context.Background()

	if _, err := repository.FindByID(ctx, "missing-user"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound for missing id, got %v", err)
	}

	if _, err := repository.FindByEmail(ctx, "missing@example.com"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound for missing email, got %v", err)
	}
}

func setupRepository(t *testing.T) *Repository {
	t.Helper()

	db := dbtest.Open(t)
	applyMigrationsIfNeeded(t, db)
	dbtest.TruncateTables(t, db, "users")

	return NewRepository(db)
}

func applyMigrationsIfNeeded(t testing.TB, db *sql.DB) {
	t.Helper()

	var hasUsersTable bool
	err := db.QueryRow(`
		SELECT EXISTS (
			SELECT 1
			FROM information_schema.tables
			WHERE table_schema = 'public'
				AND table_name = 'users'
		)
	`).Scan(&hasUsersTable)
	if err != nil {
		t.Fatalf("check users table: %v", err)
	}

	if hasUsersTable {
		return
	}

	dbtest.ApplyMigrations(t, db)
}

func createUser(t *testing.T, repository *Repository, params CreateParams) User {
	t.Helper()

	user, err := repository.Create(context.Background(), params)
	if err != nil {
		t.Fatalf("create user fixture: %v", err)
	}

	return user
}

func assertSameUser(t *testing.T, expected User, actual User) {
	t.Helper()

	if actual.ID != expected.ID {
		t.Fatalf("expected id %q, got %q", expected.ID, actual.ID)
	}
	if actual.Email != expected.Email {
		t.Fatalf("expected email %q, got %q", expected.Email, actual.Email)
	}
	if actual.Handle != expected.Handle {
		t.Fatalf("expected handle %q, got %q", expected.Handle, actual.Handle)
	}
	if actual.DisplayName != expected.DisplayName {
		t.Fatalf("expected display name %q, got %q", expected.DisplayName, actual.DisplayName)
	}
	if actual.PasswordHash != expected.PasswordHash {
		t.Fatalf("expected password hash %q, got %q", expected.PasswordHash, actual.PasswordHash)
	}
	if actual.Role != expected.Role {
		t.Fatalf("expected role %q, got %q", expected.Role, actual.Role)
	}
}
