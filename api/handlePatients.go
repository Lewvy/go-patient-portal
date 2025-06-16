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

func GetPatient(s *models.State) gin.HandlerFunc {
	return func(c *gin.Context) {
		patientName := c.Param("name")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		patient, err := s.Db.GetPatient(ctx, patientName)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"name":        patient.Name,
			"age":         patient.Age,
			"id":          patient.ID,
			"gender":      patient.Gender,
			"admitted_at": patient.CreatedAt,
			"updated_at":  patient.UpdatedAt,
		})
	}
}

func DeletePatient(s *models.State) gin.HandlerFunc {
	return func(c *gin.Context) {
		patientName := c.Param("name")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		err := s.Db.DeletePatient(ctx, patientName)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unexpected error: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Patient deleted"})
	}
}

func UpdatePatient(state *models.State) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		var req struct {
			Age       *int    `json:"age"`
			Gender    *string `json:"gender"`
			Address   *string `json:"address"`
			Diagnosis *string `json:"diagnosis"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		patient, err := state.Db.GetPatient(ctx, name)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
			return
		}

		age := patient.Age
		if req.Age != nil {
			age = int32(*req.Age)
		}
		gender := patient.Gender
		if req.Gender != nil {
			gender = *req.Gender
		}
		address := patient.Address
		if req.Address != nil {
			address = sql.NullString{String: *req.Address, Valid: *req.Address != ""}
		}

		diagnosis := patient.Diagnosis
		if req.Diagnosis != nil {
			diagnosis = sql.NullString{String: *req.Diagnosis, Valid: *req.Diagnosis != ""}
		}

		params := database.UpdatePatientDetailsParams{
			Name:      name,
			Age:       age,
			Gender:    gender,
			Address:   address,
			Diagnosis: diagnosis,
		}

		updated, err := state.Db.UpdatePatientDetails(ctx, params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed: " + err.Error()})
			return
		}
		c.JSON(http.StatusOK, updated)
	}
}
