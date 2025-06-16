package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Lewvy/markable/api"
	"github.com/Lewvy/markable/internal/config"
	"github.com/Lewvy/markable/internal/database"
	"github.com/Lewvy/markable/internal/models"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	err := config.SetConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(cfg.DB_URL)
	db, err := sql.Open("postgres", cfg.DB_URL)
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
	router.POST("/login", api.HandleLogin(state))
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
