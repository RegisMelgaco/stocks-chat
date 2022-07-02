package entity

import "github.com/google/uuid"

type Room struct {
	ID         int
	ExternalID uuid.UUID
	Name       string
}
