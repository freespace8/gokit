package cache

import (
	"time"

	"github.com/dtm-labs/rockscache"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	rc     *rockscache.Client
	prefix string
}

func NewCache(rdb redis.UniversalClient, options ...CacheOptionFunc) *Cache {
	rc := rockscache.NewClient(rdb, rockscache.NewDefaultOptions())
	c := Cache{rc: rc}
	for _, opt := range options {
		opt(&c, &rc.Options)
	}

	return &c
}

func (c *Cache) Fetch(key string, expire time.Duration, fn func() (string, error)) (string, error) {
	return c.rc.Fetch(c.prefix+key, expire, fn)
}

func (c *Cache) TagAsDeleted(key string) error {
	return c.rc.TagAsDeleted(c.prefix + key)
}
