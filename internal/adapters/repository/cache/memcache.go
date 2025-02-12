package cache

import (
	"github.com/bradfitz/gomemcache/memcache"
	"time"
)

type MemcacheClient interface {
	Set(key string, value []byte, expiration time.Duration) error
	Get(key string) ([]byte, error)
}

type memcacheClient struct {
	client *memcache.Client
}

func NewMemcacheClient() MemcacheClient {
	cacheClient := memcache.New("memcached:11211")
	err := cacheClient.Ping()
	if err != nil {
		panic(err)
	}

	return &memcacheClient{
		client: cacheClient,
	}
}

func (m *memcacheClient) Set(key string, value []byte, expiration time.Duration) error {
	return m.client.Set(&memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: int32(expiration.Seconds()),
	})
}

func (m *memcacheClient) Get(key string) ([]byte, error) {
	item, err := m.client.Get(key)
	if err != nil {
		return nil, err
	}
	return item.Value, nil
}
