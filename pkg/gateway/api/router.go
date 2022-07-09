package api

import (
	"local/stocks-chat/pkg/domain/entity"
	"local/stocks-chat/pkg/gateway/api/room"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/websocket"
)

func RouterNew(u entity.Usecases, upgrader websocket.Upgrader) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	roomHandler := room.NewHandler(u, upgrader)

	r.Get("/room", roomHandler.List)
	r.Get("/room/{room_id}/listen", roomHandler.Listen)

	return r
}
