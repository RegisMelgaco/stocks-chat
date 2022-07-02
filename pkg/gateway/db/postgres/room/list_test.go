package room

import (
	"context"
	"local/stocks-chat/pkg/domain/entity"
	"local/stocks-chat/pkg/gateway/db/postgres"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Repository_Room_List(t *testing.T) {
	t.Run("return room list when db has entries", func(t *testing.T) {
		testDB, err := postgres.CreateTestDB(testCluster)
		require.NoError(t, err)

		repo := NewRepository(testDB)

		expected := []entity.Room{
			{
				ExternalID: uuid.New(),
				ID:         1,
				Name:       "Test Room",
			},
		}

		// prepare test entry
		_, err = testDB.Exec(
			context.Background(),
			"INSERT INTO room(external_id, name) VALUES ($1, $2)",
			expected[0].ExternalID,
			expected[0].Name,
		)
		require.NoError(t, err)

		actual, err := repo.List(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}
