package entity

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	ListRooms(context.Context) ([]Room, error)
	GetRoomByExternalID(context.Context, uuid.UUID) (Room, error)
	ListMessages(ctx context.Context, filter ListMessagesFilter, limit int) ([]Message, error)
}

type ListMessagesFilter struct {
	RoomID int
}
