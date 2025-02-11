package user

import (
	"github.com/labstack/echo/v4"
	"github.com/pauloRohling/locknote/internal/application/user"
	"github.com/pauloRohling/locknote/internal/presentation/rest"
)

// RestController defines the REST API for the [user.Service]
type RestController struct {
	service       user.Service
	tokenVerifier echo.MiddlewareFunc
}

func NewRestController(service user.Service, tokenVerifier echo.MiddlewareFunc) *RestController {
	return &RestController{
		service:       service,
		tokenVerifier: tokenVerifier,
	}
}

func (controller *RestController) Register(api *echo.Group) {
	usersApi := api.Group("/users")
	usersApi.POST("", controller.create)
	usersApi.GET("", controller.get, controller.tokenVerifier)
	usersApi.POST("/login", controller.login)
}

// Ensure the controller implements the [rest.RegistrableRoute] interface
var _ rest.RegistrableRoute = (*RestController)(nil)
