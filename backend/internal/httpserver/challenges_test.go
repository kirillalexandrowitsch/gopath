package httpserver

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kirillalexandrowitsch/gopath/backend/internal/challenges"
)

func TestChallengeRunEndpoint(t *testing.T) {
	t.Parallel()

	runner := &fakeChallengeRunner{
		result: challenges.Result{
			Status:      challenges.StatusPassed,
			Stdout:      "ok",
			TestsPassed: 3,
			TestsTotal:  3,
			DurationMs:  12,
		},
	}
	request := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/challenges/retry-context/run",
		bytes.NewBufferString(`{"code":"package main"}`),
	)
	recorder := httptest.NewRecorder()

	newRouter(nil, runner).ServeHTTP(recorder, request)

	assertStatus(t, recorder, http.StatusOK)
	assertJSONContentType(t, recorder)

	var response challenges.Result
	decodeJSON(t, recorder, &response)

	if response.Status != challenges.StatusPassed {
		t.Fatalf("expected passed status, got %q", response.Status)
	}
	if runner.runCalls != 1 {
		t.Fatalf("expected 1 run call, got %d", runner.runCalls)
	}
	if runner.submitCalls != 0 {
		t.Fatalf("expected 0 submit calls, got %d", runner.submitCalls)
	}
	if runner.challengeID != "retry-context" {
		t.Fatalf("expected retry-context challenge, got %q", runner.challengeID)
	}
}

func TestChallengeSubmitEndpoint(t *testing.T) {
	t.Parallel()

	runner := &fakeChallengeRunner{
		result: challenges.Result{
			Status:      challenges.StatusPassed,
			Stdout:      "ok",
			TestsPassed: 5,
			TestsTotal:  5,
			DurationMs:  20,
		},
	}
	request := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/challenges/retry-context/submit",
		bytes.NewBufferString(`{"code":"package main"}`),
	)
	recorder := httptest.NewRecorder()

	newRouter(nil, runner).ServeHTTP(recorder, request)

	assertStatus(t, recorder, http.StatusOK)

	var response challenges.Result
	decodeJSON(t, recorder, &response)

	if response.TestsTotal != 5 {
		t.Fatalf("expected 5 tests, got %d", response.TestsTotal)
	}
	if runner.submitCalls != 1 {
		t.Fatalf("expected 1 submit call, got %d", runner.submitCalls)
	}
}

func TestChallengeInvalidJSON(t *testing.T) {
	t.Parallel()

	request := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/challenges/retry-context/run",
		bytes.NewBufferString(`{`),
	)
	recorder := httptest.NewRecorder()

	newRouter(nil, &fakeChallengeRunner{}).ServeHTTP(recorder, request)

	assertStatus(t, recorder, http.StatusBadRequest)

	var response errorResponse
	decodeJSON(t, recorder, &response)

	if response.Error != "invalid JSON" {
		t.Fatalf("expected invalid JSON error, got %q", response.Error)
	}
}

func TestChallengeEmptyCode(t *testing.T) {
	t.Parallel()

	request := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/challenges/retry-context/run",
		bytes.NewBufferString(`{"code":"   "}`),
	)
	recorder := httptest.NewRecorder()

	newRouter(nil, &fakeChallengeRunner{}).ServeHTTP(recorder, request)

	assertStatus(t, recorder, http.StatusBadRequest)

	var response errorResponse
	decodeJSON(t, recorder, &response)

	if response.Error != "code is required" {
		t.Fatalf("expected code is required error, got %q", response.Error)
	}
}

func TestChallengeNotFound(t *testing.T) {
	t.Parallel()

	request := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/challenges/unknown/run",
		bytes.NewBufferString(`{"code":"package main"}`),
	)
	recorder := httptest.NewRecorder()

	newRouter(nil, &fakeChallengeRunner{err: challenges.ErrChallengeNotFound}).ServeHTTP(recorder, request)

	assertStatus(t, recorder, http.StatusNotFound)

	var response errorResponse
	decodeJSON(t, recorder, &response)

	if response.Error != "challenge not found" {
		t.Fatalf("expected challenge not found error, got %q", response.Error)
	}
}

func TestChallengeInternalRunnerError(t *testing.T) {
	t.Parallel()

	request := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/challenges/retry-context/run",
		bytes.NewBufferString(`{"code":"package main"}`),
	)
	recorder := httptest.NewRecorder()

	newRouter(nil, &fakeChallengeRunner{err: errors.New("runner failed")}).ServeHTTP(recorder, request)

	assertStatus(t, recorder, http.StatusInternalServerError)

	var response errorResponse
	decodeJSON(t, recorder, &response)

	if response.Error != "internal server error" {
		t.Fatalf("expected internal server error, got %q", response.Error)
	}
}

type fakeChallengeRunner struct {
	result      challenges.Result
	err         error
	runCalls    int
	submitCalls int
	challengeID string
	code        string
}

func (r *fakeChallengeRunner) Run(
	ctx context.Context,
	challengeID string,
	code string,
) (challenges.Result, error) {
	r.runCalls++
	r.challengeID = challengeID
	r.code = code

	return r.result, r.err
}

func (r *fakeChallengeRunner) Submit(
	ctx context.Context,
	challengeID string,
	code string,
) (challenges.Result, error) {
	r.submitCalls++
	r.challengeID = challengeID
	r.code = code

	return r.result, r.err
}
