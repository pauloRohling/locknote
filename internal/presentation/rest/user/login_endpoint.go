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

	response, err := controller.service.Login(c.Request().Context(), &user.LoginInput{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &LoginResponse{
		AccessToken: response.AccessToken,
	})
}
