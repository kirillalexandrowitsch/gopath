//go:build integration

package attempts

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

const (
	seedLessonID    = "go-errors-tests"
	seedChallengeID = "retry-context"
)

func TestRepositoryCreateLessonAttempt(t *testing.T) {
	repositories := setupRepositories(t)
	user := createUserFixture(t, repositories.users)

	attempt, err := repositories.attempts.Create(context.Background(), CreateParams{
		ID:        "test-attempt-lesson",
		UserID:    user.ID,
		LessonID:  stringPtr(seedLessonID),
		Answer:    "caller-controls-error",
		IsCorrect: true,
		XPAwarded: 40,
		Status:    StatusPassed,
	})
	if err != nil {
		t.Fatalf("create lesson attempt: %v", err)
	}

	if attempt.ID != "test-attempt-lesson" {
		t.Fatalf("expected id %q, got %q", "test-attempt-lesson", attempt.ID)
	}
	if attempt.UserID != user.ID {
		t.Fatalf("expected user id %q, got %q", user.ID, attempt.UserID)
	}
	if attempt.LessonID == nil || *attempt.LessonID != seedLessonID {
		t.Fatalf("expected lesson id %q, got %v", seedLessonID, attempt.LessonID)
	}
	if attempt.ChallengeID != nil {
		t.Fatalf("expected nil challenge id, got %q", *attempt.ChallengeID)
	}
	if attempt.Answer != "caller-controls-error" {
		t.Fatalf("expected answer %q, got %q", "caller-controls-error", attempt.Answer)
	}
	if !attempt.IsCorrect {
		t.Fatal("expected correct attempt")
	}
	if attempt.XPAwarded != 40 {
		t.Fatalf("expected 40 XP, got %d", attempt.XPAwarded)
	}
	if attempt.Status != StatusPassed {
		t.Fatalf("expected status %q, got %q", StatusPassed, attempt.Status)
	}
	if attempt.CreatedAt.IsZero() {
		t.Fatal("expected created timestamp")
	}
}

func TestRepositoryCreateChallengeAttempt(t *testing.T) {
	repositories := setupRepositories(t)
	user := createUserFixture(t, repositories.users)

	attempt, err := repositories.attempts.Create(context.Background(), CreateParams{
		ID:          "test-attempt-challenge",
		UserID:      user.ID,
		ChallengeID: stringPtr(seedChallengeID),
		IsCorrect:   true,
		XPAwarded:   120,
		Status:      StatusPassed,
		Stdout:      "PASS",
		Stderr:      "",
		TestsPassed: 5,
		TestsTotal:  5,
		DurationMS:  812,
	})
	if err != nil {
		t.Fatalf("create challenge attempt: %v", err)
	}

	if attempt.ChallengeID == nil || *attempt.ChallengeID != seedChallengeID {
		t.Fatalf("expected challenge id %q, got %v", seedChallengeID, attempt.ChallengeID)
	}
	if attempt.LessonID != nil {
		t.Fatalf("expected nil lesson id, got %q", *attempt.LessonID)
	}
	if attempt.Stdout != "PASS" {
		t.Fatalf("expected stdout %q, got %q", "PASS", attempt.Stdout)
	}
	if attempt.TestsPassed != 5 {
		t.Fatalf("expected 5 tests passed, got %d", attempt.TestsPassed)
	}
	if attempt.TestsTotal != 5 {
		t.Fatalf("expected 5 tests total, got %d", attempt.TestsTotal)
	}
	if attempt.DurationMS != 812 {
		t.Fatalf("expected duration 812ms, got %d", attempt.DurationMS)
	}
}

func TestRepositoryFindByID(t *testing.T) {
	repositories := setupRepositories(t)
	user := createUserFixture(t, repositories.users)

	createdAttempt := createAttemptFixture(t, repositories.attempts, CreateParams{
		ID:        "test-attempt-find",
		UserID:    user.ID,
		LessonID:  stringPtr(seedLessonID),
		Answer:    "caller-controls-error",
		IsCorrect: true,
		XPAwarded: 40,
		Status:    StatusPassed,
	})

	foundAttempt, err := repositories.attempts.FindByID(context.Background(), createdAttempt.ID)
	if err != nil {
		t.Fatalf("find attempt by id: %v", err)
	}

	assertSameAttempt(t, createdAttempt, foundAttempt)
}

func TestRepositoryListByUserID(t *testing.T) {
	repositories := setupRepositories(t)
	user := createUserFixture(t, repositories.users)

	firstAttempt := createAttemptFixture(t, repositories.attempts, CreateParams{
		ID:        "test-attempt-list-a",
		UserID:    user.ID,
		LessonID:  stringPtr(seedLessonID),
		Answer:    "caller-controls-error",
		IsCorrect: true,
		XPAwarded: 40,
		Status:    StatusPassed,
	})
	time.Sleep(time.Millisecond)
	secondAttempt := createAttemptFixture(t, repositories.attempts, CreateParams{
		ID:          "test-attempt-list-b",
		UserID:      user.ID,
		ChallengeID: stringPtr(seedChallengeID),
		IsCorrect:   true,
		XPAwarded:   120,
		Status:      StatusPassed,
		TestsPassed: 5,
		TestsTotal:  5,
		DurationMS:  812,
	})

	attempts, err := repositories.attempts.ListByUserID(context.Background(), user.ID, 0)
	if err != nil {
		t.Fatalf("list attempts by user id: %v", err)
	}

	if len(attempts) != 2 {
		t.Fatalf("expected 2 attempts, got %d", len(attempts))
	}
	assertSameAttempt(t, secondAttempt, attempts[0])
	assertSameAttempt(t, firstAttempt, attempts[1])

	limitedAttempts, err := repositories.attempts.ListByUserID(context.Background(), user.ID, 1)
	if err != nil {
		t.Fatalf("list attempts by user id with limit: %v", err)
	}
	if len(limitedAttempts) != 1 {
		t.Fatalf("expected 1 limited attempt, got %d", len(limitedAttempts))
	}
	assertSameAttempt(t, secondAttempt, limitedAttempts[0])
}

func TestRepositoryFindMissingAttempt(t *testing.T) {
	repositories := setupRepositories(t)

	if _, err := repositories.attempts.FindByID(context.Background(), "missing-attempt"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound for missing attempt, got %v", err)
	}
}

type testRepositories struct {
	attempts *Repository
	users    *users.Repository
}

func setupRepositories(t *testing.T) testRepositories {
	t.Helper()

	db := dbtest.Open(t)
	applyMigrationsIfNeeded(t, db)
	dbtest.TruncateTables(t, db, "attempts", "users")

	return testRepositories{
		attempts: NewRepository(db),
		users:    users.NewRepository(db),
	}
}

func applyMigrationsIfNeeded(t testing.TB, db *sql.DB) {
	t.Helper()

	if !tableExists(t, db, "attempts") {
		dbtest.ApplyMigrations(t, db)
		return
	}
	if !seedContentExists(t, db) {
		applyMigrationFile(t, db, "000002_seed_learning_data.sql")
	}
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

func seedContentExists(t testing.TB, db *sql.DB) bool {
	t.Helper()

	var exists bool
	err := db.QueryRow(`
		SELECT EXISTS (
			SELECT 1
			FROM lessons
			WHERE id = $1
		) AND EXISTS (
			SELECT 1
			FROM challenges
			WHERE id = $2
		)
	`, seedLessonID, seedChallengeID).Scan(&exists)
	if err != nil {
		t.Fatalf("check seed content: %v", err)
	}

	return exists
}

func applyMigrationFile(t testing.TB, db *sql.DB, name string) {
	t.Helper()

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("resolve attempts test path")
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
		ID:           "test-user-attempt",
		Email:        "attempt-user@example.com",
		Handle:       "@attemptuser",
		DisplayName:  "Attempt User",
		PasswordHash: "hashed-password",
	})
	if err != nil {
		t.Fatalf("create user fixture: %v", err)
	}

	return user
}

func createAttemptFixture(t *testing.T, repository *Repository, params CreateParams) Attempt {
	t.Helper()

	attempt, err := repository.Create(context.Background(), params)
	if err != nil {
		t.Fatalf("create attempt fixture: %v", err)
	}

	return attempt
}

func assertSameAttempt(t *testing.T, expected Attempt, actual Attempt) {
	t.Helper()

	if actual.ID != expected.ID {
		t.Fatalf("expected id %q, got %q", expected.ID, actual.ID)
	}
	if actual.UserID != expected.UserID {
		t.Fatalf("expected user id %q, got %q", expected.UserID, actual.UserID)
	}
	assertStringPtr(t, "lesson id", expected.LessonID, actual.LessonID)
	assertStringPtr(t, "challenge id", expected.ChallengeID, actual.ChallengeID)
	if actual.Answer != expected.Answer {
		t.Fatalf("expected answer %q, got %q", expected.Answer, actual.Answer)
	}
	if actual.IsCorrect != expected.IsCorrect {
		t.Fatalf("expected is correct %v, got %v", expected.IsCorrect, actual.IsCorrect)
	}
	if actual.XPAwarded != expected.XPAwarded {
		t.Fatalf("expected XP awarded %d, got %d", expected.XPAwarded, actual.XPAwarded)
	}
	if actual.Status != expected.Status {
		t.Fatalf("expected status %q, got %q", expected.Status, actual.Status)
	}
	if actual.TestsPassed != expected.TestsPassed {
		t.Fatalf("expected tests passed %d, got %d", expected.TestsPassed, actual.TestsPassed)
	}
	if actual.TestsTotal != expected.TestsTotal {
		t.Fatalf("expected tests total %d, got %d", expected.TestsTotal, actual.TestsTotal)
	}
	if actual.DurationMS != expected.DurationMS {
		t.Fatalf("expected duration %d, got %d", expected.DurationMS, actual.DurationMS)
	}
}

func assertStringPtr(t *testing.T, field string, expected *string, actual *string) {
	t.Helper()

	if (actual == nil) != (expected == nil) {
		t.Fatalf("expected %s %v, got %v", field, expected, actual)
	}
	if actual != nil && *actual != *expected {
		t.Fatalf("expected %s %q, got %q", field, *expected, *actual)
	}
}

func stringPtr(value string) *string {
	return &value
}
