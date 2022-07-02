package main

import (
	"context"
	"local/stocks-chat/pkg/domain/usecases"
	"local/stocks-chat/pkg/gateway/api/rest"
	"local/stocks-chat/pkg/gateway/db/postgres/migrate"
	"local/stocks-chat/pkg/gateway/db/postgres/room"
	"net/http"
	"os"

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

	roomRepo := room.NewRepository(pool)
	u := usecases.New(roomRepo)
	r := rest.RouterNew(u)

	log.Info().Msg("listening...")

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Error().Err(err).Msg("http server failed")
	}
}
