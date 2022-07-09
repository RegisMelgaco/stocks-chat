package entity

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID         int
	ExternalID uuid.UUID
	CreatedAt  time.Time
	Author     string
	RoomID     int
	Content    string
}
