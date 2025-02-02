package user

import (
	"github.com/labstack/echo/v4"
	userApplication "github.com/pauloRohling/locknote/internal/application/user"
	"github.com/pauloRohling/locknote/internal/domain/user"
	"net/http"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	*user.User
}

func (controller *RestController) create(c echo.Context) error {
	body := new(CreateUserRequest)
	if err := c.Bind(body); err != nil {
		return err
	}

	response, err := controller.service.Create(c.Request().Context(), &userApplication.CreateUserInput{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, &CreateUserResponse{
		User: response.User,
	})
}
