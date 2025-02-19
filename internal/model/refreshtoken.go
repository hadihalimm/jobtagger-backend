package model

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	Token     uuid.UUID `db:"token" json:"token"`
	UserId    uuid.UUID `db:"user_id" json:"user_id"`
	ExpiresAt time.Time `db:"expires_at" json:"expires_at"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
