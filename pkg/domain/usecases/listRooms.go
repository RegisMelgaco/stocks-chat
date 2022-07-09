package usecases

import (
	"context"
	"local/stocks-chat/pkg/domain/entity"
	"local/stocks-chat/pkg/domain/erring"
)

func (u usecases) ListRooms(ctx context.Context) ([]entity.Room, error) {
	dErr := erring.NewWrapper("Usecases.ListRooms")

	list, err := u.repo.ListRooms(ctx)
	if err != nil {
		return nil, dErr.Wrap(err).Err()
	}

	return list, nil
}
