package entity

import "context"

type Usecases interface {
	ListRooms(context.Context) ([]Room, error)
}
