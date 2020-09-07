package route

import (
	"github.com/labstack/echo/v4"
	"github.com/mihnealun/prox/infrastructure/container"
	"net/http"
)

// PreparePublicRoutes add the necessary public routes to echo
func PreparePublicRoutes(e *echo.Echo, containerInstance container.Container) {
	for _, r := range containerInstance.GetConfig().Routes.Endpoints {
		rName := r.Name
		e.GET(r.Path, func(c echo.Context) error {
			return c.HTML(http.StatusOK, *containerInstance.GetHttpData(rName))
		})
	}
}
