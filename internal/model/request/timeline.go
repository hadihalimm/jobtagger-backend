package request

import "time"

type CreateTimeline struct {
	Content      string    `json:"content"`
	TimelineDate time.Time `json:"timeline_date"`
}

type UpdateTimeline struct {
	Content      *string    `json:"content,omitempty"`
	TimelineDate *time.Time `json:"timeline_date,omitempty"`
}
