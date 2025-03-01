package request

import "time"

type CreateInterview struct {
	Title    string    `json:"title"`
	Date     time.Time `json:"interview_date"`
	Position string    `json:"position"`
	Company  string    `json:"company"`
	Notes    string    `json:"notes"`
}

type UpdateInterview struct {
	Title    *string    `json:"title,omitempty"`
	Date     *time.Time `json:"interview_date,omitempty"`
	Position *string    `json:"position,omitempty"`
	Company  *string    `json:"company,omitempty"`
	Notes    *string    `json:"notes,omitempty"`
}
