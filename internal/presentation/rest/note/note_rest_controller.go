package note

import (
	"github.com/labstack/echo/v4"
	"github.com/pauloRohling/locknote/internal/application/note"
	"github.com/pauloRohling/locknote/internal/presentation/rest"
)

type RestController struct {
	service note.Service
}

func NewRestController(service note.Service) *RestController {
	return &RestController{service: service}
}

func (controller *RestController) Register(api *echo.Group) {
	notesApi := api.Group("/notes")
	notesApi.POST("/", controller.create)
}

// Ensure the controller implements the [rest.RegistrableRoute] interface
var _ rest.RegistrableRoute = (*RestController)(nil)
