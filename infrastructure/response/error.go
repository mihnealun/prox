package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Error used to respond with errors
type Error struct {
	Base
}

// NewError Create new Error
func NewError(code int, errors []string) (err Error) {
	err.Code = code
	err.Errors = errors
	return
}

// NewValidationError Create new ValidationErrorResponse
func NewValidationError(context echo.Context, err error) error {
	e := NewError(http.StatusBadRequest, []string{err.Error()})

	return context.JSON(e.Code, e)
}

// NewUnauthorizedError Create new UnauthorizedErrorResponse
func NewUnauthorizedError(context echo.Context, err error) error {
	e := NewError(http.StatusForbidden, []string{err.Error()})

	return context.JSON(e.Code, e)
}

// NewNotFoundError create new NotFoundResponse
func NewNotFoundError(context echo.Context, err error) error {
	e := NewError(http.StatusNotFound, []string{err.Error()})

	return context.JSON(e.Code, e)
}

// NewInternalServerError create new InternalErrorResponse
func NewInternalServerError(context echo.Context, err error) error {
	e := NewError(http.StatusInternalServerError, []string{err.Error()})

	return context.JSON(e.Code, e)
}
