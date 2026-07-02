//go:build integration

package dbtest

import (
	"testing"

	"github.com/kirillalexandrowitsch/gopath/backend/internal/database"
)

func TestOpenSkipsWithoutDatabaseURL(t *testing.T) {
	t.Setenv(database.EnvDatabaseURL, "")

	Open(t)

	t.Fatal("expected Open to skip without GOPATH_DATABASE_URL")
}
