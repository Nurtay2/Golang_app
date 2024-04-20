package ufc

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

// Fighter struct represents a UFC fighter
type Fighter struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	WeightClass string  `json:"weight_class"`
	Reach       float64 `json:"reach"`
	Wins        int     `json:"wins"`
	Losses      int     `json:"losses"`
}

// Helper function to handle errors and respond with a JSON message
func respondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// Helper function to parse fighter ID from the request URL
func getFighterIDFromRequest(r *http.Request) (int, error) {
	idParam := r.URL.Query().Get("id")
	fighterID, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, fmt.Errorf("invalid fighter ID: %v", err)
	}
	return fighterID, nil
}

// Handler function for listing all fighters
func ListFighters(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	pageStr := r.URL.Query().Get("page")
	page, _ := strconv.Atoi(pageStr)
	pageSizeStr := r.URL.Query().Get("pageSize")
	pageSize, _ := strconv.Atoi(pageSizeStr)
	filter := r.URL.Query().Get("filter")
	sortBy := r.URL.Query().Get("sortBy")
	sortOrder := r.URL.Query().Get("sortOrder")

	if page == 0 {
		page = 1
	}

	if sortBy == "" {
		sortBy = "id"
	}
	if sortOrder != "desc" {
		sortOrder = "asc"
	}

	query := "SELECT id, name, weight_class, reach, wins, losses FROM fighters"
	if filter != "" {
		_, err := strconv.Atoi(filter)
		if err == nil {
			query += " WHERE reach = " + filter
		} else {
			query += " WHERE name LIKE '%" + filter + "%' OR weight_class LIKE '%" + filter + "%'"
		}
	}
	query += " ORDER BY " + sortBy + " " + sortOrder
	if pageSize > 0 {
		query += " LIMIT " + strconv.Itoa(pageSize) + " OFFSET " + strconv.Itoa((page-1)*pageSize)
	}

	rows, err := db.Query(query)
	if err != nil {
		handleError(w, err)
		return
	}
	defer rows.Close()

	var fighters []Fighter
	for rows.Next() {
		var fighter Fighter
		err := rows.Scan(&fighter.ID, &fighter.Name, &fighter.WeightClass, &fighter.Reach, &fighter.Wins, &fighter.Losses)
		if err != nil {
			handleError(w, err)
			return
		}
		fighters = append(fighters, fighter)
	}

	writeJSONResponse(w, fighters)

}

// Handler function for creating a new fighter
func CreateFighter(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var newFighter Fighter
	err := json.NewDecoder(r.Body).Decode(&newFighter)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	insertQuery := "INSERT INTO fighters (name, weight_class, reach, wins, losses) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err = db.QueryRow(insertQuery, newFighter.Name, newFighter.WeightClass, newFighter.Reach, newFighter.Wins, newFighter.Losses).Scan(&newFighter.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating fighter")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newFighter)
}

// Handler function for getting details of a specific fighter
func GetFighter(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fighterID, err := getFighterIDFromRequest(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	fighter := getFighterDetailsFromDB(db, fighterID)
	if fighter == nil {
		respondWithError(w, http.StatusNotFound, "Fighter not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fighter)
}

// update fighter by id
func UpdateFighter(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fighterID, err := getFighterIDFromRequest(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var updatedFighter Fighter
	err = json.NewDecoder(r.Body).Decode(&updatedFighter)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	updateQuery := "UPDATE fighters SET name=$1, weight_class=$2, reach=$3, wins=$4, losses=$5 WHERE id=$6"
	_, err = db.Exec(updateQuery, updatedFighter.Name, updatedFighter.WeightClass, updatedFighter.Reach, updatedFighter.Wins, updatedFighter.Losses, fighterID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error updating fighter")
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Handler function for deleting a fighter
func DeleteFighter(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fighterID, err := getFighterIDFromRequest(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	deleteQuery := "DELETE FROM fighters WHERE id=$1"
	_, err = db.Exec(deleteQuery, fighterID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error deleting fighter")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Example function to retrieve fighters from the database
func getFightersFromDB(db *sql.DB) []Fighter {
	query := "SELECT id, name, weight, country FROM fighters"
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error querying fighters:", err)
		return nil
	}
	defer rows.Close()

	var fighters []Fighter
	for rows.Next() {
		var fighter Fighter
		err := rows.Scan(&fighter.ID, &fighter.Name, &fighter.WeightClass, &fighter.Reach, &fighter.Wins, &fighter.Losses)
		if err != nil {
			log.Println("Error scanning fighter row:", err)
			continue
		}
		fighters = append(fighters, fighter)
	}

	return fighters
}

// Example function to retrieve details of a specific fighter from the database
func getFighterDetailsFromDB(db *sql.DB, fighterID int) *Fighter {
	query := "SELECT * FROM fighters WHERE id=$1"
	row := db.QueryRow(query, fighterID)

	var fighter Fighter
	err := row.Scan(&fighter.ID, &fighter.Name, &fighter.WeightClass, &fighter.Reach, &fighter.Wins, &fighter.Losses)
	if err != nil {
		log.Println("Error scanning fighter details:", err)
		return nil
	}

	return &fighter
}

func handleError(w http.ResponseWriter, err error) {
	fmt.Println("Error:", err)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
