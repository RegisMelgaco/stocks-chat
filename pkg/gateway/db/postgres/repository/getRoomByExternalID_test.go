package repository

import (
	"context"
	"local/stocks-chat/pkg/domain/entity"
	"local/stocks-chat/pkg/gateway/db/postgres"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Repository_GetRoomByExternalID(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name        string
		expected    *entity.Room
		expectedErr error
	}

	cases := []testCase{
		{
			name: "return room when db has entries",
			expected: &entity.Room{
				ID:         1,
				ExternalID: uuid.New(),
				Name:       "Test Room",
			},
			expectedErr: nil,
		},
		{
			name:        "return room not found when db has not the entry",
			expected:    nil,
			expectedErr: entity.ErrRoomNotFound,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			testDB, err := postgres.CreateTestDB(testCluster)
			require.NoError(t, err)

			repo := NewRepository(testDB)

			if tc.expected != nil {
				// prepare test entry
				_, err = testDB.Exec(
					context.Background(),
					"INSERT INTO room(external_id, name) VALUES ($1, $2)",
					tc.expected.ExternalID,
					tc.expected.Name,
				)
				require.NoError(t, err)
			}

			testExternalID := uuid.New()
			if tc.expected != nil {
				testExternalID = tc.expected.ExternalID
			}

			actual, err := repo.GetRoomByExternalID(context.Background(), testExternalID)

			assert.ErrorIs(t, err, tc.expectedErr)

			if tc.expected != nil {
				assert.Equal(t, *tc.expected, actual)
			}
		})
	}
}
