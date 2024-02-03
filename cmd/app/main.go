package main

import (
	"fmt"
	"net/http"

	"github.com/Nurtay2/Golang_app/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	http.Handle("/", router)
	router.HandleFunc("/Fighter", handlers.GetAllFighters).Methods("GET")
	router.HandleFunc("/fighter/{id}", handlers.GetFighterByID).Methods("GET")
	router.HandleFunc("/health-check", handlers.HealthCheck).Methods("GET")

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", router)
}
