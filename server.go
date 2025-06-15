package main

import (
	"database/sql"
	"fmt"
	"log"
	// "net/http"
	//
	// "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db_url := "postgres://mark:pwd@localhost:5433/hospital?sslmode=disable"
	db, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Fatalf("Error opening db: %q", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error while pinging: %q", err)
	}
	fmt.Println("Ping successful")
}
