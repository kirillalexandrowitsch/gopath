//go:build integration

package dbtest

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/kirillalexandrowitsch/gopath/backend/internal/database"
)

const (
	openTimeout      = 5 * time.Second
	migrationTimeout = 30 * time.Second
	truncateTimeout  = 10 * time.Second
)

func Open(t testing.TB) *sql.DB {
	t.Helper()

	databaseURL := os.Getenv(database.EnvDatabaseURL)
	if databaseURL == "" {
		t.Skipf("%s is required for database integration tests", database.EnvDatabaseURL)
	}

	ctx, cancel := context.WithTimeout(context.Background(), openTimeout)
	defer cancel()

	db, err := database.Open(ctx, database.Config{URL: databaseURL})
	if err != nil {
		t.Fatalf("open integration database: %v", err)
	}

	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			t.Logf("close integration database: %v", err)
		}
	})

	return db
}

func ApplyMigrations(t testing.TB, db *sql.DB) {
	t.Helper()

	migrationFiles, err := filepath.Glob(filepath.Join(migrationsDir(t), "*.sql"))
	if err != nil {
		t.Fatalf("find migration files: %v", err)
	}
	if len(migrationFiles) == 0 {
		t.Fatal("migration files not found")
	}
	sort.Strings(migrationFiles)

	ctx, cancel := context.WithTimeout(context.Background(), migrationTimeout)
	defer cancel()

	for _, migrationFile := range migrationFiles {
		sqlBytes, err := os.ReadFile(migrationFile)
		if err != nil {
			t.Fatalf("read migration %s: %v", filepath.Base(migrationFile), err)
		}

		if _, err := db.ExecContext(ctx, string(sqlBytes)); err != nil {
			t.Fatalf("apply migration %s: %v", filepath.Base(migrationFile), err)
		}
	}
}

func TruncateTables(t testing.TB, db *sql.DB, tables ...string) {
	t.Helper()

	if len(tables) == 0 {
		return
	}

	quotedTables := make([]string, 0, len(tables))
	for _, table := range tables {
		quotedTable, err := quoteIdentifier(table)
		if err != nil {
			t.Fatalf("invalid table name %q: %v", table, err)
		}
		quotedTables = append(quotedTables, quotedTable)
	}

	ctx, cancel := context.WithTimeout(context.Background(), truncateTimeout)
	defer cancel()

	query := fmt.Sprintf("TRUNCATE TABLE %s CASCADE", strings.Join(quotedTables, ", "))
	if _, err := db.ExecContext(ctx, query); err != nil {
		t.Fatalf("truncate tables: %v", err)
	}
}

func migrationsDir(t testing.TB) string {
	t.Helper()

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("resolve dbtest source path")
	}

	return filepath.Clean(filepath.Join(filepath.Dir(filename), "..", "..", "..", "migrations"))
}

func quoteIdentifier(identifier string) (string, error) {
	if identifier == "" {
		return "", fmt.Errorf("empty identifier")
	}

	for _, r := range identifier {
		if r >= 'a' && r <= 'z' {
			continue
		}
		if r >= 'A' && r <= 'Z' {
			continue
		}
		if r >= '0' && r <= '9' {
			continue
		}
		if r == '_' {
			continue
		}

		return "", fmt.Errorf("only letters, digits and underscore are allowed")
	}

	return `"` + identifier + `"`, nil
}
