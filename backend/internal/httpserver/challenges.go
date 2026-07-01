package httpserver

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/kirillalexandrowitsch/gopath/backend/internal/challenges"
)

func (s server) challengeRunHandler(w http.ResponseWriter, r *http.Request) {
	s.handleChallenge(w, r, false)
}

func (s server) challengeSubmitHandler(w http.ResponseWriter, r *http.Request) {
	s.handleChallenge(w, r, true)
}

func (s server) handleChallenge(w http.ResponseWriter, r *http.Request, submit bool) {
	var request challenges.CodeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "invalid JSON"})
		return
	}

	if strings.TrimSpace(request.Code) == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "code is required"})
		return
	}

	challengeID := r.PathValue("id")
	result, err := runChallenge(r.Context(), s.challengeRunner, challengeID, request.Code, submit)
	if errors.Is(err, challenges.ErrChallengeNotFound) {
		writeJSON(w, http.StatusNotFound, errorResponse{Error: "challenge not found"})
		return
	}
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: "internal server error"})
		return
	}

	writeJSON(w, http.StatusOK, result)
}

func runChallenge(
	ctx context.Context,
	runner challenges.Runner,
	challengeID string,
	code string,
	submit bool,
) (challenges.Result, error) {
	if submit {
		return runner.Submit(ctx, challengeID, code)
	}

	return runner.Run(ctx, challengeID, code)
}
