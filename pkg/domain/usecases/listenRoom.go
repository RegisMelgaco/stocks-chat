package usecases

import (
	"context"
	"local/stocks-chat/pkg/domain/entity"
	"local/stocks-chat/pkg/domain/erring"

	"github.com/google/uuid"
)

func (u usecases) ListenRoom(ctx context.Context, roomID uuid.UUID, onMessage func(entity.Message) error) error {
	dErr := erring.NewWrapper("Usecases.ListenRoom")

	room, err := u.repo.GetRoomByExternalID(ctx, roomID)
	if err != nil {
		return dErr.Wrap(err).Err()
	}

	const maxMessages = 50
	filter := entity.ListMessagesFilter{RoomID: room.ID}
	list, err := u.repo.ListMessages(ctx, filter, maxMessages)
	if err != nil {
		return dErr.Wrap(err).Err()
	}

	for _, m := range list {
		onMessage(m)
	}

	err = u.messenger.ListenRoom(ctx, room.ID, onMessage)
	if err != nil {
		return dErr.Wrap(err).Err()
	}

	return nil
}
