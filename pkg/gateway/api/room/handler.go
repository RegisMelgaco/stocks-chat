package room

import (
	"local/stocks-chat/pkg/domain/entity"

	"github.com/gorilla/websocket"
)

type Handler struct {
	u        entity.Usecases
	upgrader websocket.Upgrader
}

func NewHandler(u entity.Usecases, upgrader websocket.Upgrader) Handler {
	return Handler{u, upgrader}
}
