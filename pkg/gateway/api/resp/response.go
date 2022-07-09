package resp

import (
	"context"
	"encoding/json"
	"errors"
	"local/stocks-chat/pkg/domain/entity"
	"local/stocks-chat/pkg/domain/erring"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func OK(w http.ResponseWriter, resp interface{}) {
	w.Header().Add("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Err(err).Msg("failed to write response")
	}
}

func Error(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(GetCode(err))

	if err := json.NewEncoder(w).Encode(ToErrorV1(err)); err != nil {
		log.Error().Err(err).Msg("failed to encode response error")
	}

	LogErr(r.Context(), err)
}

func GetCode(err error) int {
	if errors.Is(err, entity.ErrBadRequest) {
		return http.StatusBadRequest
	}
	if errors.Is(err, entity.ErrNotFound) {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}

func ToErrorV1(err error) ErrorV1 {
	var dErr erring.DomainErr
	if ok := errors.As(err, &dErr); ok {
		name, description := dErr.NameAndDescribe()

		return ErrorV1{
			Name:        name,
			Description: description,
		}
	}

	return ErrorV1{
		Name:        "Internal Server Error",
		Description: "The server has encountered a situation it does not know how to handle",
	}
}

func LogErr(ctx context.Context, err error) {
	requestID, ok := ctx.Value(middleware.RequestIDKey).(string)
	if !ok {
		requestID = "request id is invalid"
	}

	var level zerolog.Level
	switch GetCode(err) {
	case http.StatusBadRequest, http.StatusNotFound:
		level = zerolog.WarnLevel
	default:
		level = zerolog.ErrorLevel
	}

	var dErr erring.DomainErr
	if ok := errors.As(err, &dErr); !ok {
		log.WithLevel(level).
			Str("request_id", requestID).
			Err(err).
			Msg("error without domain error wrapper was thrown")

		return
	}

	name, description := dErr.NameAndDescribe()
	log.WithLevel(level).
		Str("name", name).
		Str("calls", dErr.Calls()).
		Str("request_id", requestID).
		Fields(map[string]interface{}{"data": dErr.GetFields()}).
		Err(dErr.InternalErr()).
		Msg(description)
}

func LogSuccess(ctx context.Context, calls string) {
	requestID, ok := ctx.Value(middleware.RequestIDKey).(string)
	if !ok {
		requestID = "request id is invalid"
	}

	log.Info().
		Str("request_id", requestID).
		Str("calls", calls).
		Msgf("handled request with success")
}
