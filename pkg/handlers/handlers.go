package handlers

import(
	"encoding/json"
	"fmt"
	"TSIS1/test"
	_"TSIS1/test"
	"net/http"
	"strconv"
	"strings"
	"github.com/gorilla/mux"

)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running n/ fighters are ready")
}

func GetFighters(http.ResponseWriter, r* http.Request){

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(testdata.Books)

	if err != nil{
		return
	}

	w.Write(jsonResponse)

}