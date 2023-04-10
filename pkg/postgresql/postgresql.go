package postgresql

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"web/internal/config"
)

func NewConnect(ctx context.Context, cfg config.StorageConfig) (*pgxpool.Pool, error) {

	connString := cfg.Url
	pool, err := pgxpool.Connect(ctx, connString)
	if err != nil {
		log.Println("failed to connect to postgres")
	}

	return pool, err
}
