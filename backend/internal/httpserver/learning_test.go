package httpserver

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
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

				var response tracksResponse
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

				var response levelsResponse
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

				var response lesson
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

				var response progressResponse
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

				var response profileResponse
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

func assertStatus(t *testing.T, recorder *httptest.ResponseRecorder, want int) {
	t.Helper()

	if recorder.Code != want {
		t.Fatalf("expected status %d, got %d", want, recorder.Code)
	}
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
