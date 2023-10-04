package cache

import (
	"time"

	"github.com/dtm-labs/rockscache"
)

type CacheOptionFunc = func(*Cache, *rockscache.Options)

func WithPrefix(prefix string) CacheOptionFunc {
	return func(c *Cache, opt *rockscache.Options) {
		c.prefix = prefix
	}
}

func WithStrongConsistency() CacheOptionFunc {
	return func(c *Cache, opt *rockscache.Options) {
		opt.StrongConsistency = true
	}
}

func WithDelay(delay time.Duration) CacheOptionFunc {
	return func(c *Cache, opt *rockscache.Options) {
		opt.Delay = delay
	}
}

func WithEmptyExpire(expire time.Duration) CacheOptionFunc {
	return func(c *Cache, opt *rockscache.Options) {
		opt.EmptyExpire = expire
	}
}

func WithLockExpire(expire time.Duration) CacheOptionFunc {
	return func(c *Cache, opt *rockscache.Options) {
		opt.LockExpire = expire
	}
}

func WithLockSleep(sleep time.Duration) CacheOptionFunc {
	return func(c *Cache, opt *rockscache.Options) {
		opt.LockSleep = sleep
	}
}

func WithRandomExpireAdjustment(adjustment float64) CacheOptionFunc {
	return func(c *Cache, opt *rockscache.Options) {
		opt.RandomExpireAdjustment = adjustment
	}
}

func WithWaitReplicasTimeout(timeout time.Duration) CacheOptionFunc {
	return func(c *Cache, opt *rockscache.Options) {
		opt.WaitReplicasTimeout = timeout
	}
}
