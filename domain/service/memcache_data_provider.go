package service

import (
	"github.com/mihnealun/prox/infrastructure/rconfig"
)

type MemcacheDataProvider struct {
	Name string
}

func NewMemcacheDataProvider(provider rconfig.DataProvider) DataProvider {
	return &MemcacheDataProvider{
		Name: provider.Name,
	}
}

func (dp *MemcacheDataProvider) GetValue(key string) string {
	return `{
	"var1": "memcache ceva"
	"var2": "memcache altCeva"
}`
}
