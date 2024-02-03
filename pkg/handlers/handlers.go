package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Nurtay2/Golang_app/api"
	"github.com/gorilla/mux"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running n/ fighters are ready")
}

func GetAllFighters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	for _, Fighter := range api.Fighters {
		json.NewEncoder(w).Encode(Fighter)
	}
}

// GetFighterByID returns a specific fighter by ID
func GetFighterByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fighterID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid fighter ID", http.StatusBadRequest)
		return
	}

	for _, fighter := range api.Fighters {
		if fighter.Id == fighterID {
			json.NewEncoder(w).Encode(fighter)
			return
		}
	}

	http.Error(w, "Fighter not found", http.StatusNotFound)
}
