package models

import (
	"github.com/google/uuid"
)

type User struct {
	UserName string    `json:"username"`
	Password string    `json:"password"`
	ID       uuid.UUID `json:"id"`
	Role     string    `json:"role"`
}
