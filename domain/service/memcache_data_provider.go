package service

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/mihnealun/prox/infrastructure/rconfig"
)

type MemcacheDataProvider struct {
	Name   string
	Client *memcache.Client
}

func NewMemcacheDataProvider(provider rconfig.DataProvider) DataProvider {
	client := memcache.New(provider.ConnectionString)

	_ = client.Set(&memcache.Item{Key: "landing", Value: []byte(`{"key": "memcache value yeeee"}`)})

	return &MemcacheDataProvider{
		Name:   provider.Name,
		Client: memcache.New(provider.ConnectionString),
	}
}

func (dp *MemcacheDataProvider) GetValue(key string) []byte {
	result, err := dp.Client.Get(key)
	if err != nil {
		return []byte{}
	}

	return result.Value
}
