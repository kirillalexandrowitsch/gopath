package challenges

import (
	"context"
	"errors"
	"os/exec"
	"testing"
)

func TestGoTestRunnerRunPublicTests(t *testing.T) {
	t.Parallel()
	skipWithoutGo(t)

	result, err := NewGoTestRunner().Run(context.Background(), "retry-context", correctRetrySolution)
	if err != nil {
		t.Fatalf("run challenge: %v", err)
	}

	if result.Status != StatusPassed {
		t.Fatalf("expected passed status, got %q\nstdout:\n%s\nstderr:\n%s", result.Status, result.Stdout, result.Stderr)
	}
	if result.TestsPassed != 3 || result.TestsTotal != 3 {
		t.Fatalf("expected 3/3 tests, got %d/%d", result.TestsPassed, result.TestsTotal)
	}
}

func TestGoTestRunnerSubmitIncludesHiddenTests(t *testing.T) {
	t.Parallel()
	skipWithoutGo(t)

	result, err := NewGoTestRunner().Submit(context.Background(), "retry-context", correctRetrySolution)
	if err != nil {
		t.Fatalf("submit challenge: %v", err)
	}

	if result.Status != StatusPassed {
		t.Fatalf("expected passed status, got %q\nstdout:\n%s\nstderr:\n%s", result.Status, result.Stdout, result.Stderr)
	}
	if result.TestsPassed != 5 || result.TestsTotal != 5 {
		t.Fatalf("expected 5/5 tests, got %d/%d", result.TestsPassed, result.TestsTotal)
	}
}

func TestGoTestRunnerChallengeNotFound(t *testing.T) {
	t.Parallel()

	_, err := NewGoTestRunner().Run(context.Background(), "unknown", correctRetrySolution)
	if !errors.Is(err, ErrChallengeNotFound) {
		t.Fatalf("expected ErrChallengeNotFound, got %v", err)
	}
}

func TestGoTestRunnerCompileError(t *testing.T) {
	t.Parallel()
	skipWithoutGo(t)

	result, err := NewGoTestRunner().Run(context.Background(), "retry-context", "package main\nfunc FetchUser(")
	if err != nil {
		t.Fatalf("run challenge: %v", err)
	}

	if result.Status != StatusCompileError {
		t.Fatalf("expected compile_error status, got %q", result.Status)
	}
	if result.TestsPassed != 0 || result.TestsTotal != 3 {
		t.Fatalf("expected 0/3 tests, got %d/%d", result.TestsPassed, result.TestsTotal)
	}
}

func skipWithoutGo(t *testing.T) {
	t.Helper()

	if _, err := exec.LookPath("go"); err != nil {
		t.Skip("go binary is not available")
	}
}

const correctRetrySolution = `package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	ID   int    ` + "`json:\"id\"`" + `
	Name string ` + "`json:\"name\"`" + `
}

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

var client Doer

func FetchUser(ctx context.Context, id int) (User, error) {
	var lastErr error
	delay := 100 * time.Millisecond

	for attempt := 0; attempt < 3; attempt++ {
		request, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://example.com/users/%d", id), nil)
		if err != nil {
			return User{}, err
		}

		response, err := client.Do(request)
		if err == nil {
			defer closeBody(response)

			if response.StatusCode != http.StatusOK {
				if response.StatusCode >= 500 {
					lastErr = fmt.Errorf("temporary status: %d", response.StatusCode)
				} else {
					return User{}, fmt.Errorf("unexpected status: %d", response.StatusCode)
				}
			} else {
				var user User
				if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
					return User{}, err
				}
				return user, nil
			}
		} else {
			if !isTemporary(err) {
				return User{}, err
			}
			lastErr = err
		}

		if attempt == 2 {
			break
		}

		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			if !timer.Stop() {
				<-timer.C
			}
			return User{}, ctx.Err()
		case <-timer.C:
		}
		delay *= 2
	}

	if lastErr == nil {
		lastErr = errors.New("fetch user failed")
	}

	return User{}, lastErr
}

func isTemporary(err error) bool {
	var temporary interface {
		Temporary() bool
	}

	return errors.As(err, &temporary) && temporary.Temporary()
}

func closeBody(response *http.Response) {
	if response != nil && response.Body != nil {
		_ = response.Body.Close()
	}
}
`
