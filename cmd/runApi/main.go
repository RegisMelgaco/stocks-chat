package main

import (
	"context"
	"local/stocks-chat/pkg/domain/usecases"
	"local/stocks-chat/pkg/gateway/api"
	"local/stocks-chat/pkg/gateway/api/room"
	"local/stocks-chat/pkg/gateway/db/postgres/migrate"
	roomRepository "local/stocks-chat/pkg/gateway/db/postgres/repository"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Error().Err(err).Msg("failed to load enviroment variables")

		return
	}

	websocketBufferSize := getIntEnvVar("WEB_SOCKET_BUFFER_SIZE")
	maxTimeWithoutHealthCheck := getIntEnvVar("WEB_SOCKET_HEALTH_CHECK_MAX_TIME")
	maxWebsocketListeners := getIntEnvVar("WEB_SOCKET_MAX_LISTENERS")
	maxWebsocketMessages := getIntEnvVar("WEB_SOCKET_MAX_MESSAGES")

	err = migrate.MigrateUP(os.Getenv("DB_URL"))
	if err != nil {
		log.Error().Err(err).Msg("failed to run migrations")

		return
	}

	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Error().Err(err).Msg("failed to create pool of connections with db")

		return
	}

	upgrader := websocket.Upgrader{
		ReadBufferSize:  websocketBufferSize,
		WriteBufferSize: websocketBufferSize,
	}
	messenger := room.NewMessenger(maxWebsocketMessages, maxWebsocketListeners, time.Duration(maxTimeWithoutHealthCheck))
	roomRepo := roomRepository.NewRepository(pool)
	u := usecases.New(roomRepo, messenger)

	r := api.RouterNew(u, upgrader)

	log.Info().Msg("listening...")

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Error().Err(err).Msg("http server failed")
	}
}

func getIntEnvVar(key string) int {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Error().Err(err).Str("key", key).Msg("env variable is invalid or not present")

		os.Exit(1)
	}

	return val
}
