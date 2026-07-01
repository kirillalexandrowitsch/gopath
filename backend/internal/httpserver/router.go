package httpserver

import (
	"net/http"

	"github.com/kirillalexandrowitsch/gopath/backend/internal/learning"
)

type server struct {
	learningStore learning.Store
}

func NewRouter() http.Handler {
	return NewRouterWithStore(learning.NewMemoryStore())
}

func NewRouterWithStore(store learning.Store) http.Handler {
	if store == nil {
		store = learning.NewMemoryStore()
	}

	server := server{
		learningStore: store,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthzHandler)
	mux.HandleFunc("GET /api/v1/tracks", server.tracksHandler)
	mux.HandleFunc("GET /api/v1/levels", server.levelsHandler)
	mux.HandleFunc("GET /api/v1/lessons/{id}", server.lessonHandler)
	mux.HandleFunc("GET /api/v1/progress", server.progressHandler)
	mux.HandleFunc("GET /api/v1/profile", server.profileHandler)

	return mux
}
