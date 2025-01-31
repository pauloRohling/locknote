package main

import (
	"fmt"
	"github.com/pauloRohling/locknote/internal/environment"
	"github.com/pauloRohling/locknote/internal/presentation/rest"
)

func main() {
	env := environment.Env()

	server := rest.NewWebServer(env.Server.Port)
	if err := server.Start(); err != nil {
		panic(fmt.Errorf("unable to start web server: %w", err))
	}
}
