package repository

import (
	"local/stocks-chat/pkg/domain/entity"

	"github.com/jackc/pgx/v4/pgxpool"
)

type repo struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) entity.Repository {
	return repo{pool}
}
