package rest

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pauloRohling/throw"
	"log/slog"
	"net/http"
	"strings"
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
	server.Use(middleware.Recover())
	server.HTTPErrorHandler = httpErrorHandler
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

func httpErrorHandler(err error, c echo.Context) {
	var throwError *throw.Error
	if !errors.As(err, &throwError) {
		slog.Error(err.Error())
		_ = c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	stringBuilder := new(strings.Builder)
	if throwError.Unwrap() == nil {
		stringBuilder.WriteString(throwError.Error())
	} else {
		stringBuilder.WriteString(fmt.Sprintf("%s: %s", throwError.Error(), throwError.Unwrap().Error()))
	}

	for _, attr := range throwError.Attributes() {
		stringBuilder.WriteString(fmt.Sprintf("%s: %s", attr.Key(), attr.Value()))
	}

	errorMessage := stringBuilder.String()
	slog.Error(errorMessage)

	_ = c.JSON(http.StatusInternalServerError, map[string]string{
		"message": errorMessage,
	})
}
