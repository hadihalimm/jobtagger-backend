package request

import "time"

type CreateJobApplication struct {
	Position    string    `json:"position"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	Source      string    `json:"source"`
	Progress    string    `json:"progress"`
	AppliedDate time.Time `json:"applied_date"`
	Notes       string    `json:"notes,omitempty"`
}

type UpdateJobApplication struct {
	Position    *string    `json:"position,omitempty"`
	Company     *string    `json:"company,omitempty"`
	Location    *string    `json:"location,omitempty"`
	Source      *string    `json:"source,omitempty"`
	Progress    *string    `json:"progress,omitempty"`
	AppliedDate *time.Time `json:"applied_date,omitempty"`
	Notes       *string    `json:"notes,omitempty"`
}
