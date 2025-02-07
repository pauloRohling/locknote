package token

import (
	"github.com/labstack/echo/v4"
	"github.com/pauloRohling/locknote/internal/domain/audit"
	"github.com/pauloRohling/locknote/internal/domain/token"
	"github.com/pauloRohling/throw"
	"strings"
)

func VerifierMiddleware(tokenVerifier token.Verifier) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenHeader := c.Request().Header.Get("Authorization")
			tokenHeader = strings.TrimPrefix(tokenHeader, "Bearer ")

			tokenPayload, err := tokenVerifier.Verify(tokenHeader)
			if err != nil {
				return throw.Unauthorized().Err(err).Msg("Missing or invalid authentication token")
			}

			ctx := audit.SetUserId(c.Request().Context(), tokenPayload.UserID)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
