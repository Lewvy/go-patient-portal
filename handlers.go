package main

import (
	"database/sql"
	"github.com/Lewvy/markable/internal/config"
	"log"
)

func loadConfig() config.Config {
	if err := config.SetConfig(); err != nil {
		log.Fatal(err)
	}
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}

func connectDB(dbURL string) *sql.DB {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening db: %q", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Error while pinging: %q", err)
	}
	log.Println("Database connection established")
	return db
}
