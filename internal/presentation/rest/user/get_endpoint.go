package user

import (
	"github.com/labstack/echo/v4"
	"github.com/pauloRohling/locknote/internal/domain/user"
	"net/http"
)

type GetResponse struct {
	*user.User
}

func (controller *RestController) get(c echo.Context) error {
	output, err := controller.service.Get(c.Request().Context())
	if err != nil {
		return err
	}

	response := &GetResponse{User: output.User}
	return c.JSON(http.StatusOK, response)
}
