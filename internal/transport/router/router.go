package router

import (
	"log"
	"net/http"

	"github.com/SussyaPusya/TZ/internal/transport/handlers"
	"github.com/gorilla/mux"
)

type Router struct {
	router *mux.Router

	handler *handlers.Handleres
}

func New(handlers *handlers.Handleres) *Router {

	r := mux.NewRouter()

	r.HandleFunc("/quotes", handlers.GetQuotes).Methods("GET")
	r.HandleFunc("/quotes/random", handlers.RandomQoute).Methods("GET")

	r.HandleFunc("/quotes", handlers.AddQuote).Methods("POST")

	r.HandleFunc("/quotes/{id}", handlers.DeleteQoute).Methods("DELETE")

	return &Router{handler: handlers, router: r}
}

func (r *Router) Run() {
	log.Fatal(http.ListenAndServe(":8080", r.router))
}
