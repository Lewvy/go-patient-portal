package main

import (
	"log"

	"github.com/Lewvy/markable/api"
	"github.com/Lewvy/markable/internal/database"
	"github.com/Lewvy/markable/internal/middleware"
	"github.com/Lewvy/markable/internal/models"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg := loadConfig()
	db := connectDB(cfg.DB_URL)
	defer db.Close()

	state := &models.State{
		Db:  database.New(db),
		Cfg: &cfg,
	}

	router := setupRouter(state)
	log.Println("Server listening on :8888")
	if err := router.Run(":8888"); err != nil {
		log.Fatalf("%q", err)
	}
}

func setupRouter(state *models.State) *gin.Engine {
	router := gin.Default()
	router.GET("/ping", handleRoot)
	router.POST("/register", api.HandleRegistration(state))
	router.POST("/login", api.HandleLogin(state))

	patients := router.Group("/patients")
	patients.Use(middleware.AuthorizeToken(state))

	patients.POST("", middleware.RequiredRole("Receptionist"), api.CreatePatient(state))
	patients.GET("/:name", middleware.RequiredRole("Receptionist", "Doctor"), api.GetPatient(state))
	patients.DELETE("/:name", middleware.RequiredRole("Receptionist"), api.DeletePatient(state))
	patients.PATCH("/:name", middleware.RequiredRole("Receptionist", "Doctor"), api.UpdatePatient(state))

	return router
}

func handleRoot(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
