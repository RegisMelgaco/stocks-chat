package entity

import "context"

type RoomRepository interface {
	List(context.Context) ([]Room, error)
}
