package user

import (
	"github.com/labstack/echo/v4"
	"github.com/pauloRohling/locknote/internal/application/user"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

func (controller *RestController) login(c echo.Context) error {
	body := new(LoginRequest)
	if err := c.Bind(body); err != nil {
		return err
	}

	input := &user.LoginInput{
		Email:    body.Email,
		Password: body.Password,
	}

	output, err := controller.service.Login(c.Request().Context(), input)
	if err != nil {
		return err
	}

	response := &LoginResponse{AccessToken: output.AccessToken}
	return c.JSON(http.StatusOK, response)
}
