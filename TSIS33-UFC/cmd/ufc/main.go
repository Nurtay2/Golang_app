package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Nurtay2/TSIS2-UFC/pkg/ufc"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "Nurtay2006@"
	dbname   = "ufcfighters"
)

func initDB() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the UFC database")
	return db
}

func main() {
	db := initDB()
	defer db.Close()

	http.HandleFunc("/fighters", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ufc.ListFighters(w, r, db)
		case http.MethodPost:
			ufc.CreateFighter(w, r, db)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/fighters/", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/fighters/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid fighter ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			r.URL.RawQuery = "id=" + strconv.Itoa(id)
			ufc.GetFighter(w, r, db)
		case http.MethodPut:
			r.URL.RawQuery = "id=" + strconv.Itoa(id)
			ufc.UpdateFighter(w, r, db)
		case http.MethodDelete:
			r.URL.RawQuery = "id=" + strconv.Itoa(id)
			ufc.DeleteFighter(w, r, db)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	port := "8080"
	log.Printf("Server running on :%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
