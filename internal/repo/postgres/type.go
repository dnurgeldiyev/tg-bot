package postgres

import (
	"dovran/tg-bot/internal/repo"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Postgres struct {
	pg *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) repo.Repo {
	return &Postgres{pg: pool}
}
