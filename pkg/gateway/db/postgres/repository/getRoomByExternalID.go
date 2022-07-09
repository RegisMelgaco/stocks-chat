package repository

import (
	"context"
	"local/stocks-chat/pkg/domain/entity"
	"local/stocks-chat/pkg/domain/erring"

	"github.com/google/uuid"
)

func (r repo) GetRoomByExternalID(ctx context.Context, id uuid.UUID) (entity.Room, error) {
	dErr := erring.NewWrapper("Repository.GetRoomByExternalID")
	dErr.Field("room.external_id", id)

	const query = "SELECT id, name FROM room WHERE external_id = $1 LIMIT 1"

	room := entity.Room{ExternalID: id}
	if err := r.pool.QueryRow(ctx, query, id).Scan(&room.ID, &room.Name); err != nil {
		err = entity.ErrRoomNotFound.Wrap(err)

		return entity.Room{}, dErr.Wrap(err).Err()
	}

	return room, nil
}
