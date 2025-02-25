package model

import (
	"time"

	"github.com/google/uuid"
)

type JobApplication struct {
	ID          int       `db:"id" json:"id"`
	UserID      uuid.UUID `db:"user_id" json:"user_id"`
	Position    string    `db:"position" json:"position"`
	Company     string    `db:"company" json:"company"`
	Location    string    `db:"location" json:"location"`
	Source      string    `db:"source" json:"source"`
	Progress    string    `db:"progress" json:"progress"`
	AppliedDate time.Time `db:"applied_date" json:"applied_date"`
	Notes       string    `db:"notes" json:"notes"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
