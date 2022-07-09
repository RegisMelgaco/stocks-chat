package entity

import (
	"context"
)

type Messenger interface {
	ListenRoom(ctx context.Context, roomID int, onMessage func(Message) error) error
	SendMessage(ctx context.Context, p Message)
}
