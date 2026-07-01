package httpserver

import "net/http"

type healthResponse struct {
	Status string `json:"status"`
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, healthResponse{Status: "ok"})
}
