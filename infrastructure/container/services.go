package container

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mihnealun/prox/domain/service"
	"github.com/mihnealun/prox/infrastructure/rconfig"
	"io/ioutil"
	"sync"
	"time"
)

// Container interface that described what services it holds
type Container interface {
	GetConfig() *Config
	GetLogger(ctx context.Context) (Logger, error)
	GetHttpData(routeName string) *string
	SetHttpData(routeName string, content string)
	SetProvider(name string, provider service.DataProvider)
	GetProvider(name string) (service.DataProvider, error)
	AddRoute(path, name string)
	GetRouteNameByPath(path string) (string, error)
	InitRefreshingStaticData() error
	RegisterProviders() error
	BuildPage(template, data []byte) string
}

type container struct {
	config    *Config
	Html      map[string]*string
	providers map[string]service.DataProvider
	templates map[string]string
	routes    map[string]string
}

var instance *container
var once sync.Once

// GetInstance return the container as a singleton instance
func GetInstance() (c Container, err error) {
	once.Do(func() {
		instance = &container{}

		instance.config, err = getConfigInstance()

		instance.Html = make(map[string]*string)
		instance.providers = make(map[string]service.DataProvider)
		instance.templates = make(map[string]string)
		instance.routes = make(map[string]string)

		if err != nil {
			return
		}
	})

	return instance, err
}

// GetConfig is returning the Config instance
func (c *container) GetConfig() *Config {
	return c.config
}

// SetProvider
func (c *container) SetProvider(name string, provider service.DataProvider) {
	if _, ok := c.providers[name]; ok {
		return
	}

	c.providers[name] = provider
}

// AddRoute
func (c *container) AddRoute(path, name string) {
	if _, ok := c.routes[name]; ok {
		return
	}
	c.routes[path] = name
}

// GetProvider
func (c *container) GetRouteNameByPath(path string) (string, error) {
	if routeName, ok := c.routes[path]; ok {
		return routeName, nil
	}

	return "", fmt.Errorf("route %s is not registered", path)
}

// GetProvider
func (c *container) GetProvider(name string) (service.DataProvider, error) {
	if providerService, ok := c.providers[name]; ok {
		return providerService, nil
	}

	return nil, fmt.Errorf("provider %s is not implemented", name)
}

// GetLogger provides application logger (can be customize if required for third party integration)
func (c *container) GetLogger(ctx context.Context) (Logger, error) {
	return newStdLogger(ctx, c.config)
}

func getRouteConfig() (data rconfig.Config, err error) {
	file, err := ioutil.ReadFile("routes.json")
	if err != nil {
		return data, err
	}

	err = json.Unmarshal([]byte(file), &data)

	if err != nil {
		return data, err
	}

	return data, nil
}

// SetHttpData updates the http static content for a route
func (c *container) SetHttpData(routeName, content string) {
	c.Html[routeName] = &content
}

// GetHttpData returns the static HTTP data for a given route
func (c *container) GetHttpData(routeName string) *string {
	if data, ok := c.Html[routeName]; ok {
		return data
	}

	emptyResult := ""

	return &emptyResult
}

// InitRefreshingStaticData rebuilds and updates the httpdata every N seconds
func (c *container) InitRefreshingStaticData() error {
	for _, ep := range c.GetConfig().Routes.Endpoints {
		c.AddRoute(ep.Path, ep.Name)
		provider, err := c.GetProvider(ep.Data.Provider)
		if err != nil {
			return err
		}

		templateProvider, err := c.GetProvider(ep.Template.Provider)
		if err != nil {
			return err
		}

		go func(c Container, cep rconfig.Endpoint, prov service.DataProvider, templateProv service.DataProvider) {
			for {
				data := prov.GetValue(cep.Data.Key)
				template := templateProv.GetValue(cep.Template.Key)
				c.SetHttpData(cep.Name, c.BuildPage(template, data))
				time.Sleep(time.Second * time.Duration(cep.Data.Ttl))
			}
		}(c, ep, provider, templateProvider)
	}

	return nil
}

func (c *container) BuildPage(template, data []byte) string {
	// TODO: implement templating service and merge data into template
	return string(data)
}

// RegisterProviders register the data providers so that they can be used by routes
func (c *container) RegisterProviders() error {
	for _, dp := range c.GetConfig().Routes.Providers {
		if dp.Name == "memcache" {
			c.SetProvider(dp.Name, service.NewMemcacheDataProvider(dp))
			continue
		}

		if dp.Name == "redis" {
			c.SetProvider(dp.Name, service.NewRedisDataProvider(dp))
			continue
		}

		return fmt.Errorf("provider not implemented")
	}

	return nil
}
