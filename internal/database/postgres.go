package database

import (
	"context"
	"kasir-api/internal/config"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/rs/zerolog/log"
)

func InitPostgresDB(ctx context.Context, config *config.Config) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(config.DBURL)
	if err != nil {
		return nil, err
	}

	cfg.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger: tracelog.LoggerFunc(func(
			ctx context.Context,
			level tracelog.LogLevel,
			msg string,
			data map[string]any,
		) {
			switch level {
			case tracelog.LogLevelDebug:
				log.Debug().Fields(data).Msg(msg)
			case tracelog.LogLevelInfo:
				log.Info().Fields(data).Msg(msg)
			case tracelog.LogLevelWarn:
				log.Warn().Fields(data).Msg(msg)
			case tracelog.LogLevelError:
				log.Error().Fields(data).Msg(msg)
			}
		}),
		LogLevel: tracelog.LogLevelInfo,
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
