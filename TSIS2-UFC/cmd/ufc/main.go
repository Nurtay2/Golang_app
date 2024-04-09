package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

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
		ufc.ListFighters(w, r, db)
	})
	http.HandleFunc("/fighters/create", func(w http.ResponseWriter, r *http.Request) {
		ufc.CreateFighter(w, r, db)
	})
	http.HandleFunc("/fighters/get", func(w http.ResponseWriter, r *http.Request) {
		ufc.GetFighter(w, r, db)
	})
	http.HandleFunc("/fighters/update", func(w http.ResponseWriter, r *http.Request) {
		ufc.UpdateFighter(w, r, db)
	})
	http.HandleFunc("/fighters/delete", func(w http.ResponseWriter, r *http.Request) {
		ufc.DeleteFighter(w, r, db)
	})

	port := "8080"
	log.Printf("UFC Server running on :%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
