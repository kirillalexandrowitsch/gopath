package challenges

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const (
	runnerTimeout = 2 * time.Second
)

var passedTestPattern = regexp.MustCompile(`(?m)^--- PASS:\s+([A-Za-z0-9_]+)`)

type GoTestRunner struct{}

func NewGoTestRunner() *GoTestRunner {
	return &GoTestRunner{}
}

func (r *GoTestRunner) Run(ctx context.Context, challengeID string, code string) (Result, error) {
	return r.execute(ctx, challengeID, code, false)
}

func (r *GoTestRunner) Submit(ctx context.Context, challengeID string, code string) (Result, error) {
	return r.execute(ctx, challengeID, code, true)
}

func (r *GoTestRunner) execute(ctx context.Context, challengeID string, code string, includeHidden bool) (Result, error) {
	spec, ok := lookupChallenge(challengeID)
	if !ok {
		return Result{}, ErrChallengeNotFound
	}

	tests := append([]challengeTest(nil), spec.PublicTests...)
	if includeHidden {
		tests = append(tests, spec.HiddenTests...)
	}

	startedAt := time.Now()
	ctx, cancel := context.WithTimeout(ctx, runnerTimeout)
	defer cancel()

	tempDir, err := os.MkdirTemp("", "gopath-challenge-*")
	if err != nil {
		return Result{}, fmt.Errorf("create challenge temp dir: %w", err)
	}
	defer os.RemoveAll(tempDir)

	if err := writeChallengeFiles(tempDir, code, tests); err != nil {
		return Result{}, err
	}

	stdout, stderr, err := runCommand(ctx, tempDir, "gofmt", "-w", "solution.go")
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		return timeoutResult(startedAt, stdout, stderr, len(tests)), nil
	}
	if err != nil {
		if isExitError(err) {
			return Result{
				Status:      StatusCompileError,
				Stdout:      stdout,
				Stderr:      stderr,
				TestsPassed: 0,
				TestsTotal:  len(tests),
				DurationMs:  elapsedMillis(startedAt),
				Hint:        "Проверьте синтаксис Go-кода.",
			}, nil
		}

		return Result{}, fmt.Errorf("run gofmt: %w", err)
	}

	stdout, stderr, err = runCommand(ctx, tempDir, "go", "test", "-count=1", "-timeout=2s", "-v", "./...")
	if errors.Is(ctx.Err(), context.DeadlineExceeded) || isGoTestTimeout(stdout, stderr) {
		return timeoutResult(startedAt, stdout, stderr, len(tests)), nil
	}
	if err != nil {
		status := StatusFailed
		hint := "Проверьте логи тестов и исправьте поведение функции."
		if isCompileFailure(stdout, stderr) {
			status = StatusCompileError
			hint = "Проверьте ошибки компиляции и сигнатуру функции."
		}

		return Result{
			Status:      status,
			Stdout:      stdout,
			Stderr:      stderr,
			TestsPassed: countPassedTests(stdout),
			TestsTotal:  len(tests),
			DurationMs:  elapsedMillis(startedAt),
			Hint:        hint,
		}, nil
	}

	return Result{
		Status:      StatusPassed,
		Stdout:      stdout,
		Stderr:      stderr,
		TestsPassed: len(tests),
		TestsTotal:  len(tests),
		DurationMs:  elapsedMillis(startedAt),
		Hint:        "",
	}, nil
}

func writeChallengeFiles(dir string, code string, tests []challengeTest) error {
	files := map[string]string{
		"go.mod":           "module challenge\n\ngo 1.26\n",
		"solution.go":      code,
		"solution_test.go": buildTestFile(tests),
	}

	for name, content := range files {
		path := filepath.Join(dir, name)
		if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
			return fmt.Errorf("write %s: %w", name, err)
		}
	}

	return nil
}

func runCommand(ctx context.Context, dir string, name string, args ...string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	command := exec.CommandContext(ctx, name, args...)
	command.Dir = dir
	command.Stdout = &stdout
	command.Stderr = &stderr

	err := command.Run()

	return stdout.String(), stderr.String(), err
}

func timeoutResult(startedAt time.Time, stdout string, stderr string, testsTotal int) Result {
	return Result{
		Status:      StatusTimeout,
		Stdout:      stdout,
		Stderr:      stderr,
		TestsPassed: countPassedTests(stdout),
		TestsTotal:  testsTotal,
		DurationMs:  elapsedMillis(startedAt),
		Hint:        "Проверьте бесконечные циклы, блокировки и обработку context.Context.",
	}
}

func elapsedMillis(startedAt time.Time) int64 {
	return time.Since(startedAt).Milliseconds()
}

func isExitError(err error) bool {
	var exitError *exec.ExitError
	return errors.As(err, &exitError)
}

func countPassedTests(stdout string) int {
	matches := passedTestPattern.FindAllStringSubmatch(stdout, -1)
	return len(matches)
}

func isCompileFailure(stdout string, stderr string) bool {
	output := stdout + "\n" + stderr

	return strings.Contains(output, "[build failed]") ||
		strings.Contains(output, "undefined:") ||
		strings.Contains(output, "expected") ||
		strings.Contains(output, "syntax error")
}

func isGoTestTimeout(stdout string, stderr string) bool {
	output := stdout + "\n" + stderr

	return strings.Contains(output, "panic: test timed out")
}
