package user

import (
	"github.com/labstack/echo/v4"
	"github.com/pauloRohling/locknote/internal/application/user"
	"github.com/pauloRohling/locknote/internal/presentation/rest"
)

type RestController struct {
	service user.Service
}

func NewRestController(service user.Service) *RestController {
	return &RestController{service: service}
}

func (controller *RestController) Register(api *echo.Group) {
	usersApi := api.Group("/users")
	usersApi.POST("", controller.create)
	usersApi.POST("/login", controller.login)
}

// Ensure the controller implements the [rest.RegistrableRoute] interface
var _ rest.RegistrableRoute = (*RestController)(nil)
