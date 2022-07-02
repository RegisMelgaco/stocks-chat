package rest

import (
	"local/stocks-chat/pkg/domain/entity"
	"local/stocks-chat/pkg/gateway/api/rest/room"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RouterNew(u entity.Usecases) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	roomHandler := room.NewHandler(u)

	r.Get("/room", roomHandler.List)

	return r
}
