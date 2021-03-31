package webserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

func defineRoutes(router *mux.Router) {
	router.HandleFunc("/", get).Methods(http.MethodGet)
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}
