package database

import (
	"context"
	"kasir-api/internal/config"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitPostgresDB(ctx context.Context, config *config.Config) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(config.DBURL)
	if err != nil {
		return nil, err
	}

	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.MaxConnIdleTime = 5 * time.Minute
	cfg.MaxConnLifetime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}
