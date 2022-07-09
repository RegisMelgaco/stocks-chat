package room

import (
	"context"
	"local/stocks-chat/pkg/domain/entity"
	"local/stocks-chat/pkg/domain/erring"
	"local/stocks-chat/pkg/gateway/api/resp"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func (h Handler) Listen(w http.ResponseWriter, r *http.Request) {
	dErr := erring.NewWrapper("Handler.Room.Listen")

	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		resp.LogErr(
			r.Context(),
			dErr.Wrap(err).Err(),
		)

		return
	}

	roomID, err := uuid.Parse(chi.URLParam(r, "room_id"))
	if err != nil {
		err = dErr.Wrap(
			entity.ErrInvalidID.Wrap(err),
		).Err()

		_ = conn.WriteJSON(resp.ToErrorV1(err))

		resp.LogErr(r.Context(), err)

		if err := conn.Close(); err != nil {
			resp.LogErr(r.Context(), err)
		}

		return
	}

	err = h.u.ListenRoom(r.Context(), roomID, onMessage(conn))
	if err != nil {
		err = dErr.Wrap(err).Err()

		_ = conn.WriteJSON(resp.ToErrorV1(err))

		resp.LogErr(r.Context(), err)

		if err := conn.Close(); err != nil {
			resp.LogErr(r.Context(), err)
		}
	}

	resp.LogSuccess(r.Context(), dErr.Calls())
}

func onMessage(conn *websocket.Conn) func(entity.Message) error {
	return func(m entity.Message) error {
		dErr := erring.NewWrapper("Handler.Room.Listen.onMessage")

		err := conn.WriteJSON(ToMessageOutput(m))
		if err != nil {
			resp.LogErr(
				context.Background(),
				dErr.Wrap(err).Err(),
			)

			return err
		}

		return nil
	}
}
