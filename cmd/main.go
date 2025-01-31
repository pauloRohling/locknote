package main

import (
	"context"
	"fmt"
	"github.com/pauloRohling/locknote/internal/environment"
	"github.com/pauloRohling/locknote/internal/persistence/postgres"
	"github.com/pauloRohling/locknote/internal/presentation/rest"
	"log/slog"
)

func main() {
	env := environment.Env()

	dbPoolBuilder := postgres.NewPoolBuilder(env.GetDatabaseAddress(), env.GetDatabaseUrl())
	dbPool := dbPoolBuilder.Build(context.Background())
	defer dbPool.Close()

	slog.Info("Database connection established")

	server := rest.NewWebServer(env.Server.Port)
	if err := server.Start(); err != nil {
		panic(fmt.Errorf("unable to start web server: %w", err))
	}
}
