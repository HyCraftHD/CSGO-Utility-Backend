package webserver

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	log.Println("Start webserver")

	router := mux.NewRouter()
	defineRoutes(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
