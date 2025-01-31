package rest

import (
	"fmt"
	"github.com/labstack/echo/v4"
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
	server := echo.New()
	server.HideBanner = true

	api := server.Group("/api/v1")

	return &WebServer{
		port:   port,
		server: server,
		api:    api,
	}
}

func (server *WebServer) Start() error {
	address := fmt.Sprintf(":%d", server.port)
	return server.server.Start(address)
}

func (server *WebServer) Register(Registrable RegistrableRoute) {
	Registrable.Register(server.api)
}
