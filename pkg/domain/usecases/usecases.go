package usecases

import "local/stocks-chat/pkg/domain/entity"

type usecases struct {
	roomRepo entity.RoomRepository
}

func New(roomRepo entity.RoomRepository) entity.Usecases {
	return usecases{roomRepo}
}
