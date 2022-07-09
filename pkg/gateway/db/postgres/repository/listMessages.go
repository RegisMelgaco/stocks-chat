package repository

import (
	"context"
	"local/stocks-chat/pkg/domain/entity"
	"local/stocks-chat/pkg/domain/erring"
)

func (r repo) ListMessages(ctx context.Context, filter entity.ListMessagesFilter, limit int) ([]entity.Message, error) {
	dErr := erring.NewWrapper("Repository.ListMessages")

	const sql = `
		SELECT
			id, external_id, content, author, created_at, room_id
		FROM 
			message
		WHERE 
			room_id = $1
		ORDER BY
			created_at DESC
		LIMIT 
			$2
	`

	rows, err := r.pool.Query(ctx, sql, filter.RoomID, limit)
	if err != nil {
		return nil, dErr.Wrap(err).Err()
	}

	list := make([]entity.Message, 0)
	for rows.Next() {
		var m entity.Message
		err := rows.Scan(&m.ID, &m.ExternalID, &m.Content, &m.Author, &m.CreatedAt, &m.RoomID)
		if err != nil {
			return nil, dErr.Wrap(err).Err()
		}

		list = append(list, m)
	}

	return list, nil
}
