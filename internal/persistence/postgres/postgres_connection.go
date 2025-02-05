package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PoolBuilder struct {
	address string
	url     string
	log     *zap.Logger
}

func NewPoolBuilder(address, url string, log *zap.Logger) *PoolBuilder {
	return &PoolBuilder{
		address: address,
		url:     url,
		log:     log,
	}
}

func (builder *PoolBuilder) Build(ctx context.Context) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(builder.url)
	if err != nil {
		builder.log.Fatal("Unable to create database pool configuration", zap.Error(err))
	}

	dbPool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		builder.log.Fatal("Unable to connect to database", zap.Error(err))
	}

	if err = dbPool.Ping(ctx); err != nil {
		builder.log.Fatal("Unable to ping database", zap.Error(err))
	}

	return dbPool
}
