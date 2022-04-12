package handler

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RequestHandler() {
	r := mux.NewRouter()
	r.HandleFunc("/link", PasteLink).Methods("POST")
	r.HandleFunc("/link", GetLink).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
