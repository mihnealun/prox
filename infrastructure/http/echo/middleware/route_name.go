package middleware

import (
	"github.com/labstack/echo/v4"

	"github.com/mihnealun/prox/infrastructure/container"
)

// RouteName is a middleware where you can translate path to routeName
func RouteName(c container.Container) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			r, err := c.GetRouteNameByPath(ctx.Path())
			if err != nil {
				return err
			}

			ctx.Set("routeName", r)

			return next(ctx)
		}
	}
}
