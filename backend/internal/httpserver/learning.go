package httpserver

import (
	"errors"
	"net/http"

	"github.com/kirillalexandrowitsch/gopath/backend/internal/learning"
)

const demoUserID = "demo-user"

func (s server) tracksHandler(w http.ResponseWriter, r *http.Request) {
	tracks, err := s.learningStore.Tracks(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, errorInternalServer)
		return
	}

	writeJSON(w, http.StatusOK, tracks)
}

func (s server) levelsHandler(w http.ResponseWriter, r *http.Request) {
	levels, err := s.learningStore.Levels(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, errorInternalServer)
		return
	}

	writeJSON(w, http.StatusOK, levels)
}

func (s server) lessonHandler(w http.ResponseWriter, r *http.Request) {
	lessonID := r.PathValue("id")
	lesson, err := s.learningStore.Lesson(r.Context(), lessonID)
	if errors.Is(err, learning.ErrLessonNotFound) {
		writeError(w, http.StatusNotFound, errorLessonNotFound)
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, errorInternalServer)
		return
	}

	writeJSON(w, http.StatusOK, lesson)
}

func (s server) progressHandler(w http.ResponseWriter, r *http.Request) {
	progress, err := s.learningStore.Progress(r.Context(), demoUserID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, errorInternalServer)
		return
	}

	writeJSON(w, http.StatusOK, progress)
}

func (s server) profileHandler(w http.ResponseWriter, r *http.Request) {
	profile, err := s.learningStore.Profile(r.Context(), demoUserID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, errorInternalServer)
		return
	}

	writeJSON(w, http.StatusOK, profile)
}
