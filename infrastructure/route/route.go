package route

import (
	"github.com/labstack/echo/v4"
	"github.com/mihnealun/prox/domain/service"
	"github.com/mihnealun/prox/infrastructure/rconfig"
	"net/http"
	"time"

	"github.com/mihnealun/prox/infrastructure/container"
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

func InitRefreshingStaticData(containerInstance container.Container) error {
	for _, ep := range containerInstance.GetConfig().Routes.Endpoints {
		containerInstance.AddRoute(ep.Path, ep.Name)
		provider, err := containerInstance.GetProvider(ep.Data.Provider)
		if err != nil {
			return err
		}

		go func(c container.Container, cep rconfig.Endpoint, prov service.DataProvider) {
			for {
				data := prov.GetValue(cep.Data.Key)
				//template := "asdf asdfasdf"
				c.SetHttpData(cep.Name, &data)
				time.Sleep(time.Second * time.Duration(cep.Data.Ttl))
			}
		}(containerInstance, ep, provider)
	}

	return nil
}

func RegisterProviders(containerInstance container.Container, providers []rconfig.DataProvider) {
	for _, dp := range providers {
		if dp.Name == "memcache" {
			containerInstance.SetProvider(dp.Name, service.NewMemcacheDataProvider(dp))
			continue
		}

		if dp.Name == "redis" {
			containerInstance.SetProvider(dp.Name, service.NewRedisDataProvider(dp))
		}
	}
}
