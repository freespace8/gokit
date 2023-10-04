# cache
### How to use it.
```
import (
  "github.com/freespace8/cache"
  "github.com/redis/go-redis/v9"
)

var rdb = redis.NewClient(&redis.Options{
  Addr: "localhost:6379",
  Username: "root",
  Password: "",
})

rc := cache.NewCache(rdb,
  cache.WithPrefix("cache:"),
  cache.WithDelay(0*time.Second),
)

rc.Fetch("key", 30*time.Second, func() (string, error) {
  time.Sleep(1*time.Second)
  return "value", nil
})
```
