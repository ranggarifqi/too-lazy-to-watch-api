package user

import (
	"time"

	"github.com/google/uuid"
)

var TABLE_NAME = "Users"

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
