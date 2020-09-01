package handler

import (
	"errors"

	"github.com/mihnealun/prox/infrastructure/response"

	"net/http"

	"github.com/labstack/echo/v4"
)

// HTTPErrorHandler is the HTTP error handler. It sends a JSON response
// with status code.
func HTTPErrorHandler(err error, c echo.Context) {
	var (
		code      = http.StatusInternalServerError
		msg       interface{}
		echoError *echo.HTTPError
	)

	switch {
	case errors.As(err, &echoError):
		code = echoError.Code
		msg = echoError.Message
	case c.Echo().Debug:
		msg = err.Error()
	default:
		msg = http.StatusText(code)
	}

	if _, ok := msg.(string); ok {
		msg = response.NewError(code, []string{msg.(string)})
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead { // Issue #608
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, msg)
		}

		if err != nil {
			c.Echo().Logger.Error(err)
		}
	}
}
