package app

import (
	"context"

	"dovran/tg-bot/config"
	"dovran/tg-bot/internal/repo/postgres"
	"dovran/tg-bot/internal/usecase"
	"dovran/tg-bot/pkg/logger"
	"dovran/tg-bot/pkg/postrges"
)

func Start(cfg *config.Config) {
	l := logger.New("warn")

	pg, err := postrges.NewClient(context.Background(), cfg)
	if err != nil {
		l.Fatal(err)
	}

	repo := postgres.NewRepository(pg)

	_ = usecase.NewUcase(repo)

}
