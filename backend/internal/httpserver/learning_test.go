package httpserver

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kirillalexandrowitsch/gopath/backend/internal/learning"
)

func TestCoreLearningEndpoints(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		path       string
		assertBody func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "tracks",
			path: "/api/v1/tracks",
			assertBody: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				t.Helper()

				var response learning.TracksResponse
				decodeJSON(t, recorder, &response)

				if len(response.Tracks) != 1 {
					t.Fatalf("expected 1 track, got %d", len(response.Tracks))
				}
				if response.Tracks[0].ID != "go-backend" {
					t.Fatalf("expected go-backend track, got %q", response.Tracks[0].ID)
				}
			},
		},
		{
			name: "levels",
			path: "/api/v1/levels",
			assertBody: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				t.Helper()

				var response learning.LevelsResponse
				decodeJSON(t, recorder, &response)

				if len(response.Levels) != 4 {
					t.Fatalf("expected 4 levels, got %d", len(response.Levels))
				}
				if response.Levels[1].ID != "junior-backend" {
					t.Fatalf("expected junior-backend level, got %q", response.Levels[1].ID)
				}
			},
		},
		{
			name: "lesson by id",
			path: "/api/v1/lessons/go-errors-tests",
			assertBody: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				t.Helper()

				var response learning.Lesson
				decodeJSON(t, recorder, &response)

				if response.ID != "go-errors-tests" {
					t.Fatalf("expected lesson id go-errors-tests, got %q", response.ID)
				}
				if response.Question.CorrectID != "caller-controls-error" {
					t.Fatalf("expected correct answer caller-controls-error, got %q", response.Question.CorrectID)
				}
			},
		},
		{
			name: "progress",
			path: "/api/v1/progress",
			assertBody: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				t.Helper()

				var response learning.Progress
				decodeJSON(t, recorder, &response)

				if response.XP != 1420 {
					t.Fatalf("expected 1420 XP, got %d", response.XP)
				}
				if response.CurrentLessonID != "go-errors-tests" {
					t.Fatalf("expected current lesson go-errors-tests, got %q", response.CurrentLessonID)
				}
			},
		},
		{
			name: "profile",
			path: "/api/v1/profile",
			assertBody: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				t.Helper()

				var response learning.Profile
				decodeJSON(t, recorder, &response)

				if response.User.Name != "Alex Kim" {
					t.Fatalf("expected profile for Alex Kim, got %q", response.User.Name)
				}
				if response.Readiness.TargetLevel != "Middle Backend" {
					t.Fatalf("expected Middle Backend target, got %q", response.Readiness.TargetLevel)
				}
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			request := httptest.NewRequest(http.MethodGet, test.path, nil)
			recorder := httptest.NewRecorder()

			NewRouter().ServeHTTP(recorder, request)

			assertStatus(t, recorder, http.StatusOK)
			assertJSONContentType(t, recorder)
			test.assertBody(t, recorder)
		})
	}
}

func TestLessonNotFound(t *testing.T) {
	t.Parallel()

	request := httptest.NewRequest(http.MethodGet, "/api/v1/lessons/unknown-lesson", nil)
	recorder := httptest.NewRecorder()

	NewRouter().ServeHTTP(recorder, request)

	assertStatus(t, recorder, http.StatusNotFound)
	assertJSONContentType(t, recorder)

	var response errorResponse
	decodeJSON(t, recorder, &response)

	if response.Error != "lesson not found" {
		t.Fatalf("expected lesson not found error, got %q", response.Error)
	}
}

func TestLearningStoreErrorReturnsInternalServerError(t *testing.T) {
	t.Parallel()

	request := httptest.NewRequest(http.MethodGet, "/api/v1/tracks", nil)
	recorder := httptest.NewRecorder()

	NewRouterWithStore(failingLearningStore{}).ServeHTTP(recorder, request)

	assertStatus(t, recorder, http.StatusInternalServerError)
	assertJSONContentType(t, recorder)

	var response errorResponse
	decodeJSON(t, recorder, &response)

	if response.Error != "internal server error" {
		t.Fatalf("expected internal server error, got %q", response.Error)
	}
}

func assertStatus(t *testing.T, recorder *httptest.ResponseRecorder, want int) {
	t.Helper()

	if recorder.Code != want {
		t.Fatalf("expected status %d, got %d", want, recorder.Code)
	}
}

type failingLearningStore struct{}

func (s failingLearningStore) Tracks(ctx context.Context) (learning.TracksResponse, error) {
	return learning.TracksResponse{}, errors.New("store failed")
}

func (s failingLearningStore) Levels(ctx context.Context) (learning.LevelsResponse, error) {
	return learning.LevelsResponse{}, errors.New("store failed")
}

func (s failingLearningStore) Lesson(ctx context.Context, id string) (learning.Lesson, error) {
	return learning.Lesson{}, errors.New("store failed")
}

func (s failingLearningStore) Progress(ctx context.Context, userID string) (learning.Progress, error) {
	return learning.Progress{}, errors.New("store failed")
}

func (s failingLearningStore) Profile(ctx context.Context, userID string) (learning.Profile, error) {
	return learning.Profile{}, errors.New("store failed")
}

func assertJSONContentType(t *testing.T, recorder *httptest.ResponseRecorder) {
	t.Helper()

	if contentType := recorder.Header().Get("Content-Type"); contentType != "application/json" {
		t.Fatalf("expected content type application/json, got %q", contentType)
	}
}

func decodeJSON(t *testing.T, recorder *httptest.ResponseRecorder, destination any) {
	t.Helper()

	if err := json.NewDecoder(recorder.Body).Decode(destination); err != nil {
		t.Fatalf("decode response: %v", err)
	}
}
