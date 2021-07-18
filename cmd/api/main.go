package main

import (
	"context"
	"database/sql"
	"flag"
	"os"
	"time"

	"github.com/adamszpilewicz/rest-temperature/internal/data"
	"github.com/adamszpilewicz/rest-temperature/internal/jsonlog"
	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port string
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	logger *jsonlog.Logger
	models data.Models
}

func main() {
	var cfg config

	flag.StringVar(&cfg.port, "port", "8080", "port to serve the application")
	flag.StringVar(&cfg.env, "env", "DEV", "environment serving the application")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://pi:kub16ala@localhost/pi", "PostgreSQL DSN")
	flag.Parse()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	db, err := openDB(cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	defer db.Close()

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	app.logger.PrintInfo("success: db connection pool established", nil)


	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
