package error

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/pauloRohling/locknote/pkg/array"
	"github.com/pauloRohling/throw"
	"go.uber.org/zap"
	"net/http"
)

func NewErrorHandler(log *zap.Logger) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if err == nil {
			return
		}

		var customError *throw.Error
		if !errors.As(err, &customError) {
			customError = throw.Internal().Err(err).Msgf("Unexpected error")
		}

		log.Error(err.Error(), array.Map(customError.Attributes(), func(attr throw.Attribute) zap.Field {
			return zap.String(attr.Key(), attr.Value())
		})...)

		statusCode := throw.ErrorType(customError.Type()).StatusCode()
		_ = c.JSON(statusCode, HTTPError{
			Err:    customError.Unwrap(),
			Title:  http.StatusText(statusCode),
			Detail: customError.Error(),
			Status: statusCode,
		})
	}
}
