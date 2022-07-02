package room

import "local/stocks-chat/pkg/domain/entity"

type Handler struct {
	u entity.Usecases
}

func NewHandler(u entity.Usecases) Handler {
	return Handler{u}
}
