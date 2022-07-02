package room

import "local/stocks-chat/pkg/domain/entity"

type RoomOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ToOutput(rooms []entity.Room) []RoomOutput {
	out := make([]RoomOutput, 0, len(rooms))
	for _, r := range rooms {
		out = append(out, RoomOutput{
			ID:   r.ExternalID.String(),
			Name: r.Name,
		})
	}

	return out
}
