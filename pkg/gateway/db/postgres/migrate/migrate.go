package migrate

import (
	"context"
	"embed"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	migratePGX "github.com/golang-migrate/migrate/v4/database/pgx"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/zerolog/log"
)

//go:embed *.sql
var fs embed.FS

func MigrateUP(connStr string) error {
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return err
	}
	d, err := iofs.New(fs, ".")
	if err != nil {
		return err
	}

	stdConn := stdlib.OpenDB(*conn.Config())
	mDriver, err := migratePGX.WithInstance(stdConn, &migratePGX.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("iofs", d, "testDB", mDriver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && errors.Is(err, migrate.ErrNoChange) {
		log.Warn().Err(err).Msg("no migration to run")
	}

	return nil
}
