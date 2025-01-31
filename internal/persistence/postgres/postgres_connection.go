package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PoolBuilder struct {
	address string
	url     string
}

func NewPoolBuilder(address, url string) *PoolBuilder {
	return &PoolBuilder{
		address: address,
		url:     url,
	}
}

func (builder *PoolBuilder) Build(ctx context.Context) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(builder.url)
	if err != nil {
		panic("Unable to create database pool configuration")
	}

	dbPool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		panic("Unable to connect to database")
	}

	if err = dbPool.Ping(ctx); err != nil {
		panic("Unable to ping database")
	}

	return dbPool
}
