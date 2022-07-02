package room

import (
	"local/stocks-chat/pkg/domain/erring"
	"local/stocks-chat/pkg/gateway/api/rest/resp"
	"net/http"
)

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	dErr := erring.New("Handler.Room.List")

	list, err := h.u.ListRooms(r.Context())
	if err != nil {
		resp.Error(w, dErr.Wrap(err).Err())

		return
	}

	resp.OK(w, ToOutput(list))

	return
}
