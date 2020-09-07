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
	result := &MemcacheDataProvider{
		Name:   provider.Name,
		Client: memcache.New(provider.ConnectionString),
	}

	// TODO: remove this after testing is done
	result.addTestData(provider)

	return result
}

func (dp *MemcacheDataProvider) GetValue(key string) []byte {
	result, err := dp.Client.Get(key)
	if err != nil {
		return []byte{}
	}

	return result.Value
}

// TODO: remove this after testing is done
func (dp *MemcacheDataProvider) addTestData(provider rconfig.DataProvider) {
	client := memcache.New(provider.ConnectionString)
	_ = client.Set(&memcache.Item{Key: "landing", Value: []byte(`{"name": "Bigus Dickus"}`)})
	_ = client.Set(&memcache.Item{Key: "landing_template", Value: []byte(`Template from memcache says your name is {{ name }}!`)})
}
