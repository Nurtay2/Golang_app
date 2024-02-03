package main

import (
	"log"
	"net/http"
	"TSIS1/pkg/handlers"
	_"TSIS1/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/health-check", handlers.HealthCheck()).Methods("GET")
	router.HandleFunc("/fighters", handlers.GetFighters()).Methods("GET")
	http.Handle("/", router)

	log.Print("starting server on :8080")
	http.ListenAndServe(":8080", router)

}
