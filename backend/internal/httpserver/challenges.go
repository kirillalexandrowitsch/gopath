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
		writeError(w, http.StatusBadRequest, errorInvalidJSON)
		return
	}

	if strings.TrimSpace(request.Code) == "" {
		writeError(w, http.StatusBadRequest, errorCodeRequired)
		return
	}

	challengeID := r.PathValue("id")
	result, err := runChallenge(r.Context(), s.challengeRunner, challengeID, request.Code, submit)
	if errors.Is(err, challenges.ErrChallengeNotFound) {
		writeError(w, http.StatusNotFound, errorChallengeNotFound)
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, errorInternalServer)
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
