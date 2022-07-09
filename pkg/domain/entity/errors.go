package entity

import (
	"local/stocks-chat/pkg/domain/erring"
)

var (
	// base
	ErrBadRequest = erring.NewErr("bad request", "request is incorrect")
	ErrNotFound   = erring.NewErr("not found", "desired entry was not found")

	// internal
	ErrMaxListenersReached = erring.NewErr("max listeners reached", "application has reached max safe listeners set by operators")

	// specific
	ErrInvalidID    = erring.NewErr("invalid id", "receaved id has invalid format or value").Wrap(ErrBadRequest)
	ErrRoomNotFound = erring.NewErr("room not found", "no room was found with specified id").Wrap(ErrNotFound)
)
