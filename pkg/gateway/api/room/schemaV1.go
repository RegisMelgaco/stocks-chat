package room

import (
	"local/stocks-chat/pkg/domain/entity"
	"time"

	"github.com/google/uuid"
)

type RoomOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MessageOutput struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
}

func ToMessageOutput(m entity.Message) MessageOutput {
	return MessageOutput{
		ID:        m.ExternalID,
		CreatedAt: m.CreatedAt,
		Author:    m.Author,
		Content:   m.Content,
	}
}

func ToRoomsOutput(rooms []entity.Room) []RoomOutput {
	out := make([]RoomOutput, 0, len(rooms))
	for _, r := range rooms {
		out = append(out, RoomOutput{
			ID:   r.ExternalID.String(),
			Name: r.Name,
		})
	}

	return out
}
