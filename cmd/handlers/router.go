package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() {
	r := mux.NewRouter()
	// r.HandleFunc("/", HomeHandler)
	// r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
}
