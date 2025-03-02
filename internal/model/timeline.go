package model

import "time"

type Timeline struct {
	ID            int       `db:"id" json:"id"`
	ApplicationID int       `db:"application_id" json:"application_id"`
	Content       string    `db:"content" json:"content"`
	TimelineDate  time.Time `db:"timeline_date" json:"timeline_date"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}
