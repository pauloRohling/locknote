package rest

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pauloRohling/locknote/internal/environment"
)

type RegistrableRoute interface {
	Register(api *echo.Group)
}

type WebServer struct {
	port   int
	server *echo.Echo
	api    *echo.Group
}

func NewWebServer(port int) *WebServer {
	env := environment.Env()

	server := echo.New()
	server.Use(middleware.Recover())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     env.Security.Cors.AllowedOrigins,
		AllowMethods:     env.Security.Cors.AllowedMethods,
		AllowHeaders:     env.Security.Cors.AllowedHeaders,
		AllowCredentials: *env.Security.Cors.AllowCredentials,
	}))

	server.HidePort = true
	server.HideBanner = true

	return &WebServer{
		port:   port,
		server: server,
		api:    server.Group("/api/v1"),
	}
}

func (server *WebServer) Start() error {
	address := fmt.Sprintf(":%d", server.port)
	return server.server.Start(address)
}

func (server *WebServer) Shutdown(ctx context.Context) error {
	return server.server.Shutdown(ctx)
}

func (server *WebServer) Register(Registrable RegistrableRoute) {
	Registrable.Register(server.api)
}

func (server *WebServer) SetErrorHandler(errorHandler echo.HTTPErrorHandler) {
	server.server.HTTPErrorHandler = errorHandler
}
