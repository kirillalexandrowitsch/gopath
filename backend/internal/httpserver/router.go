package httpserver

import (
	"net/http"

	"github.com/kirillalexandrowitsch/gopath/backend/internal/challenges"
	"github.com/kirillalexandrowitsch/gopath/backend/internal/learning"
)

type server struct {
	learningStore   learning.Store
	challengeRunner challenges.Runner
}

func NewRouter() http.Handler {
	return NewRouterWithStore(learning.NewMemoryStore())
}

func NewRouterWithStore(store learning.Store) http.Handler {
	return newRouter(store, challenges.NewGoTestRunner())
}

func newRouter(store learning.Store, runner challenges.Runner) http.Handler {
	if store == nil {
		store = learning.NewMemoryStore()
	}
	if runner == nil {
		runner = challenges.NewGoTestRunner()
	}

	server := server{
		learningStore:   store,
		challengeRunner: runner,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthzHandler)
	mux.HandleFunc("GET /api/v1/tracks", server.tracksHandler)
	mux.HandleFunc("GET /api/v1/levels", server.levelsHandler)
	mux.HandleFunc("GET /api/v1/lessons/{id}", server.lessonHandler)
	mux.HandleFunc("GET /api/v1/progress", server.progressHandler)
	mux.HandleFunc("GET /api/v1/profile", server.profileHandler)
	mux.HandleFunc("POST /api/v1/challenges/{id}/run", server.challengeRunHandler)
	mux.HandleFunc("POST /api/v1/challenges/{id}/submit", server.challengeSubmitHandler)

	return mux
}
