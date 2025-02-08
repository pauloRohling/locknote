package pagination

import (
	"github.com/labstack/echo/v4"
	"github.com/pauloRohling/locknote/internal/domain/pagination"
	"github.com/pauloRohling/throw"
	"net/url"
	"strconv"
)

func Middleware() echo.MiddlewareFunc {
	return middleware
}

func middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		queryParams := c.QueryParams()

		page, err := getQueryParam(queryParams, "page", 1)
		if err != nil {
			return throw.Validation().Err(err).Msg("Page is not a valid integer")
		}

		size, err := getQueryParam(queryParams, "size", 10)
		if err != nil {
			return throw.Validation().Err(err).Msg("Size is not a valid integer")
		}

		params := pagination.NewPagination(page, size)
		ctx := SetPagination(c.Request().Context(), params)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

func getQueryParam(queryParams url.Values, key string, defaultValue int32) (int32, error) {
	if !queryParams.Has(key) {
		return defaultValue, nil
	}

	param := queryParams.Get(key)
	paramValue, err := strconv.ParseInt(param, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(paramValue), nil
}
