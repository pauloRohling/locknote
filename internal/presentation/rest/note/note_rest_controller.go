package note

import (
	"github.com/labstack/echo/v4"
	"github.com/pauloRohling/locknote/internal/application/note"
	"github.com/pauloRohling/locknote/internal/domain/audit"
	"github.com/pauloRohling/locknote/internal/domain/token"
	"github.com/pauloRohling/locknote/internal/presentation/rest"
	"strings"
)

type RestController struct {
	service       note.Service
	tokenVerifier token.Verifier
}

func NewRestController(service note.Service, tokenVerifier token.Verifier) *RestController {
	return &RestController{
		service:       service,
		tokenVerifier: tokenVerifier,
	}
}

func (controller *RestController) Register(api *echo.Group) {
	notesApi := api.Group("/notes")

	notesApi.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenHeader := c.Request().Header.Get("Authorization")
			tokenHeader = strings.TrimPrefix(tokenHeader, "Bearer ")

			tokenPayload, err := controller.tokenVerifier.Verify(tokenHeader)
			if err != nil {
				return err
			}

			ctx := audit.SetUserId(c.Request().Context(), tokenPayload.UserID)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	})

	notesApi.POST("", controller.create)
}

// Ensure the controller implements the [rest.RegistrableRoute] interface
var _ rest.RegistrableRoute = (*RestController)(nil)
