package usecases

import (
	"context"
	"local/stocks-chat/pkg/domain/entity"
	"local/stocks-chat/pkg/domain/erring"
)

func (u usecases) ListRooms(ctx context.Context) ([]entity.Room, error) {
	dErr := erring.New("Usecases.ListRooms")

	list, err := u.roomRepo.List(ctx)
	if err != nil {
		return nil, dErr.Wrap(err).Err()
	}

	return list, nil
}
