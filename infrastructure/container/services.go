package container

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mihnealun/prox/domain/service"
	"github.com/mihnealun/prox/infrastructure/rconfig"
	"io/ioutil"
	"sync"
)

// Container interface that described what services it holds
type Container interface {
	GetConfig() *Config
	GetLogger(ctx context.Context) (Logger, error)
	GetHttpData(routeName string) *string
	SetHttpData(routeName string, content *string)
	SetProvider(name string, provider service.DataProvider)
	GetProvider(name string) (service.DataProvider, error)
	AddRoute(path, name string)
	GetRouteNameByPath(path string) (string, error)
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
func (c *container) SetHttpData(routeName string, content *string) {
	c.Html[routeName] = content
}

// GetHttpData returns the static HTTP data for a given route
func (c *container) GetHttpData(routeName string) *string {
	return c.Html[routeName]
}
