package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Lewvy/markable/api"
	"github.com/Lewvy/markable/internal/database"
	"github.com/Lewvy/markable/internal/models"

	"github.com/gin-gonic/gin"
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
	dbQueries := database.New(db)
	state := &models.State{
		Db: dbQueries,
	}
	router := gin.Default()

	router.GET("/ping", handleRoot)
	router.POST("/register", api.HandleRegistration(state))
	fmt.Println("Server listening on :8888")
	err = router.Run()
	if err != nil {
		log.Fatalf("%q", err)
	}

}

func handleRoot(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
