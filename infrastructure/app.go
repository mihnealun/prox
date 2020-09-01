package infrastructure

import (
	"fmt"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/mihnealun/prox/infrastructure/container"
	"github.com/mihnealun/prox/infrastructure/http/echo/handler"
	"github.com/mihnealun/prox/infrastructure/route"
)

// Start method is bootstrapping and starting the entire application
func Start(containerInstance container.Container) error {
	e := echo.New()
	config := containerInstance.GetConfig()

	e.HTTPErrorHandler = handler.HTTPErrorHandler

	// necessary for managing the go panic errors
	e.Use(echoMiddleware.Recover())
	//e.Use(middleware.Logger(containerInstance))
	//e.Use(middleware.RouteName(containerInstance))
	// enable http compression
	e.Use(echoMiddleware.Gzip())

	// register data providers
	route.RegisterProviders(containerInstance, containerInstance.GetConfig().Routes.Providers)

	// init workers that will regenerate content every ttl seconds for each endpoint
	err := route.InitRefreshingStaticData(containerInstance)
	if err != nil {
		return err
	}

	route.PreparePublicRoutes(e, containerInstance)

	address := fmt.Sprintf("%s:%d", config.Interface, config.Port)
	err = e.Start(address)
	if err != nil {
		return err
	}

	return nil
}
