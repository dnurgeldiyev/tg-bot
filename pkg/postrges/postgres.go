package postrges

import (
	"context"
	"fmt"

	"dovran/tg-bot/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewClient(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.Storage.Username, cfg.Storage.Password, cfg.Storage.Host, cfg.Storage.Port, cfg.Storage.Database)

	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
