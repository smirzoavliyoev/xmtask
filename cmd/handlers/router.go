package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(handlers *Handlers) {
	r := mux.NewRouter()
	r.HandleFunc("/companies", handlers.GetCompany).Methods("GET")
	r.HandleFunc("/company", handlers.Create).Methods("POST")
	r.HandleFunc("/company", handlers.Update).Methods("PUT")
	r.HandleFunc("/company", handlers.Delete).Methods("DELETE")
	http.Handle("/", r)
}
