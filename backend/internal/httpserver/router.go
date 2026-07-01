package httpserver

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthzHandler)
	mux.HandleFunc("GET /api/v1/tracks", tracksHandler)
	mux.HandleFunc("GET /api/v1/levels", levelsHandler)
	mux.HandleFunc("GET /api/v1/lessons/{id}", lessonHandler)
	mux.HandleFunc("GET /api/v1/progress", progressHandler)
	mux.HandleFunc("GET /api/v1/profile", profileHandler)

	return mux
}
