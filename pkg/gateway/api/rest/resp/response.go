package resp

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func OK(w http.ResponseWriter, resp interface{}) {
	w.Header().Add("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Err(err).Msg("failed to write response")
	}
}

func Error(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"msg": "server internal error"}`))

	log.Error().Err(err).Msg("failed to process request")
}
