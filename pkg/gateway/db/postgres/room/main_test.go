package room

import (
	"local/stocks-chat/pkg/gateway/db/postgres"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v4"
)

var testCluster *pgx.Conn

func TestMain(m *testing.M) {
	var (
		cleanup func()
		err     error
	)
	testCluster, cleanup, err = postgres.CreateTestCluster()
	if err != nil {
		log.Fatalf("failed to create test db: %s", err.Error())

		return
	}

	code := m.Run()

	cleanup()

	os.Exit(code)
}
