package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserName  string    `json:"username"`
	Password  string    `json:"-"`
	ID        uuid.UUID `json:"id"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
