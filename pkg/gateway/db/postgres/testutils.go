package postgres

import (
	"context"
	"fmt"
	"local/stocks-chat/pkg/gateway/db/postgres/migrate"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

func CreateTestCluster() (conn *pgx.Conn, cleanup func(), err error) {
	dPool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := dPool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14",
		Env: []string{
			"POSTGRES_PASSWORD=secret",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	connStr := fmt.Sprintf("postgresql://postgres:secret@localhost:%s/postgres", resource.GetPort("5432/tcp"))

	if err := dPool.Retry(func() error {
		var err error
		conn, err = pgx.Connect(context.Background(), connStr)
		if err != nil {
			return err
		}

		return conn.Ping(context.Background())
	}); err != nil {
		return nil, nil, err
	}

	cleanup = func() {
		if err := dPool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}

	return conn, cleanup, nil
}

func CreateTestDB(conn *pgx.Conn) (*pgxpool.Pool, error) {
	dbName := "test_db_" + strings.ReplaceAll(uuid.NewString(), "-", "")
	_, err := conn.Exec(context.Background(), `CREATE DATABASE `+dbName)
	if err != nil {
		log.Fatalf("failed to run create database command: CREATE DATABASE %s", dbName)

		return nil, err
	}

	connStr := fmt.Sprintf("postgresql://postgres:secret@localhost:%v/%s", conn.Config().Port, dbName)

	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	migrate.MigrateUP(connStr)

	return pool, nil
}
