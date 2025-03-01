package model

import (
	"time"

	"github.com/google/uuid"
)

type Contact struct {
	ID        int       `db:"id" json:"id"`
	UserID    uuid.UUID `db:"user_id" json:"user_id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Phone     string    `db:"phone" json:"phone"`
	Notes     string    `db:"notes" json:"notes"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
