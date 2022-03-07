package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ccontreraso/backend_app/models"
	_ "github.com/lib/pq"
)

const version = "0.0.1"

type config struct {
	port int
	env string
	db struct {
		dsn string
	}
}

type AppStatus struct {
	Status string `json:"status"`
	Environment string `json:"environment"`
	Version string `json:"version"`
}

type application struct {
	config config
	logger *log.Logger
	models models.Models
}

func main (){
	var cfg config

	flag.IntVar(&cfg.port, "port", 3005, "port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment(development|production)")
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://cristobalcontreras@localhost:5432/go_movies?sslmode=disable", "database dsn")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	app := &application{
		config: cfg,
		logger: logger,
		models: models.NewModels(db),
	}

	src := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting server on port %d", cfg.port)
	
	err = src.ListenAndServe()
	if err != nil {
		log.Println("Error:", err)
	}
};

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel  := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}