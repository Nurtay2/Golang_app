package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/Nurtay2/Golang_app/TSIS1/api"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello to my phone book!\n")
	fmt.Fprintf(w, "You can use route /people to find all people in this phonebook.\n")
	fmt.Fprintf(w, "You can use route /people/{id} to find person by ID.\n")
	fmt.Fprintf(w, "You can use route /people/name/{name} to find person by name.\n")
	fmt.Fprintf(w, "You can use route /health-check to find some info about this phonebook.\n")
}

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
	for _, Fighter := range api.Fighters {
		if Fighter.Id == params["id"]{
			json.NewEncoder(w).Encode(Fighter)
			return
		}
	}
}

