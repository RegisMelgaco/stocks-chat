package repository

import (
	"context"
	"local/stocks-chat/pkg/domain/entity"
	"local/stocks-chat/pkg/domain/erring"
)

func (r repo) ListRooms(ctx context.Context) ([]entity.Room, error) {
	dErr := erring.NewWrapper("Repository.Room.List")

	rows, err := r.pool.Query(ctx, "SELECT id, external_id, name FROM room")
	if err != nil {
		return nil, dErr.Wrap(err).Err()
	}

	list := make([]entity.Room, 0)
	for rows.Next() {
		var room entity.Room
		err := rows.Scan(&room.ID, &room.ExternalID, &room.Name)
		if err != nil {
			return nil, dErr.Wrap(err).Err()
		}

		list = append(list, room)
	}

	return list, nil
}
