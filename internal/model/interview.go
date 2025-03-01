package model

import "time"

type Interview struct {
	ID            int       `db:"id" json:"id"`
	ApplicationID int       `db:"application_id" json:"application_id"`
	Title         string    `db:"title" json:"interview_title"`
	Date          time.Time `db:"interview_date" json:"interview_date"`
	Position      string    `db:"position" json:"position"`
	Company       string    `db:"company" json:"company"`
	Notes         string    `db:"notes" json:"notes"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}
