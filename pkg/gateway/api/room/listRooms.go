package room

import (
	"local/stocks-chat/pkg/domain/erring"
	"local/stocks-chat/pkg/gateway/api/resp"
	"net/http"
)

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	dErr := erring.NewWrapper("Handler.Room.List")

	list, err := h.u.ListRooms(r.Context())
	if err != nil {
		resp.Error(w, r, dErr.Wrap(err).Err())

		return
	}

	resp.OK(w, ToRoomsOutput(list))

	return
}
