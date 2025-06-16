package api

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/Lewvy/markable/internal/database"
	"github.com/Lewvy/markable/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreatePatient(state *models.State) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Name    string `json:"name" binding:"required"`
			Age     int    `json:"age" binding:"required"`
			Gender  string `json:"gender" binding:"required"`
			Address string `json:"address"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		params := database.CreatePatientParams{
			ID:        uuid.New(),
			Name:      req.Name,
			Gender:    req.Gender,
			Age:       int32(req.Age),
			Address:   sql.NullString{Valid: false},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if req.Address != "" {
			params.Address = sql.NullString{String: req.Address, Valid: true}
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		patient, err := state.Db.CreatePatient(ctx, params)
		if err != nil {
			log.Print(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create patient"})
		}
		c.JSON(http.StatusOK, patient)
	}
}
