package error

import (
	"errors"
	"fmt"
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

		c.Response().Header().Set("Content-Type", "application/problem+json")

		var customError *throw.Error
		if !errors.As(err, &customError) {
			customError = throw.Internal().Err(err).Msgf("Unexpected error")
		}

		instance := c.Request().URL.Path
		requestID := c.Response().Header().Get(echo.HeaderXRequestID)
		statusCode := throw.ErrorType(customError.Type()).StatusCode()
		statusText := http.StatusText(statusCode)

		logError(customError, log.With(
			zap.String("requestId", requestID),
			zap.String("instance", instance),
			zap.String("method", c.Request().Method),
		))

		_ = c.JSON(statusCode, HTTPError{
			Title:     statusText,
			Status:    statusCode,
			Detail:    customError.Error(),
			Instance:  instance,
			RequestID: requestID,
		})
	}
}

func logError(customError *throw.Error, log *zap.Logger) {
	errorMessage := customError.Error()
	if innerError := customError.UnwrapOriginal(); innerError != nil {
		errorMessage = fmt.Sprintf("%s: %s", customError, innerError)
	}

	fields := array.Map(customError.Attributes(), func(attr throw.Attribute) zap.Field {
		return zap.String(attr.Key(), attr.Value())
	})

	log.Error(errorMessage, fields...)
}
