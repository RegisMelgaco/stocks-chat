package entity

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Usecases interface {
	ListRooms(context.Context) ([]Room, error)
	// SendMessage(context.Context, SendMessageInput) (Message, error)
	ListenRoom(ctx context.Context, roomID uuid.UUID, onMessage func(Message) error) error
}

type SendMessageInput struct {
	CreatedAt      time.Time
	Author         string
	RoomExternalID uuid.UUID
	Content        string
}
