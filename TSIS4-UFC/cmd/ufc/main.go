package main

import (
	"database/sql"
	"flag"
	"os"
	"sync"

	"github.com/Nurtay2/TSIS2-UFC/pkg/ufc/model"
	"github.com/Nurtay2/TSIS2-UFC/pkg/jsonlog"
	"github.com/Nurtay2/TSIS2-UFC/pkg/vcs"



	_ "github.com/lib/pq"
)

// Set version of application corresponding to value of vcs.Version.
var (
	version = vcs.Version()
)

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	models model.Models
	logger *jsonlog.Logger
	wg     sync.WaitGroup
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://postgres:Nurtay2006@@localhost:5432/demoufcfighters?sslmode=disable", "PostgreSQL DSN")
	flag.Parse()

	logger := jsonlog.NewLogger(os.Stdout, jsonlog.LevelInfo)

	db, err := openDB(cfg)
	if err != nil {
		logger.PrintError(err, nil)
		return
	}

	defer func() {
		if err := db.Close(); err != nil {
			logger.PrintFatal(err, nil)
		}
	}()

	app := &application{
		config: cfg,
		models: model.NewModels(db),
		logger: logger,
	}

	if err := app.serve(); err != nil {
		logger.PrintFatal(err, nil)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}