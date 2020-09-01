package service

import (
	"github.com/mihnealun/prox/infrastructure/rconfig"
)

type RedisDataProvider struct {
	Name string
}

func NewRedisDataProvider(provider rconfig.DataProvider) DataProvider {
	return &RedisDataProvider{
		Name: provider.Name,
	}
}

func (dp *RedisDataProvider) GetValue(key string) string {
	return `{
	"var1": "redis ceva"
	"var2": "redis altCeva"
}`
}
