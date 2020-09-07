package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/mihnealun/prox/infrastructure/rconfig"
)

type RedisDataProvider struct {
	Name   string
	Client *redis.Client
}

func NewRedisDataProvider(provider rconfig.DataProvider) DataProvider {
	return &RedisDataProvider{
		Name: provider.Name,
		Client: redis.NewClient(&redis.Options{
			Addr:     provider.ConnectionString,
			Password: provider.Password,
			DB:       0,
		}),
	}
}

func (dp *RedisDataProvider) GetValue(key string) []byte {
	result, err := dp.Client.Get(context.Background(), key).Result()
	if err != nil {
		return []byte("")
	}

	return []byte(result)
}
