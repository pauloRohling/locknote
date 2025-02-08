package note

import (
	"github.com/labstack/echo/v4"
	"github.com/pauloRohling/locknote/internal/application/note"
	"github.com/pauloRohling/locknote/internal/presentation/rest"
	"github.com/pauloRohling/locknote/internal/presentation/rest/pagination"
)

// RestController defines the REST API for the [note.Service]
type RestController struct {
	service       note.Service
	tokenVerifier echo.MiddlewareFunc
}

func NewRestController(service note.Service, tokenVerifier echo.MiddlewareFunc) *RestController {
	return &RestController{
		service:       service,
		tokenVerifier: tokenVerifier,
	}
}

func (controller *RestController) Register(api *echo.Group) {
	notesApi := api.Group("/notes")
	notesApi.Use(controller.tokenVerifier)
	notesApi.POST("", controller.create)
	notesApi.GET("", controller.list, pagination.Middleware())
	notesApi.GET("/:id", controller.getById)
}

// Ensure the controller implements the [rest.RegistrableRoute] interface
var _ rest.RegistrableRoute = (*RestController)(nil)
