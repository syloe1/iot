package redisx

import (
	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func NewClient(c redis.RedisConf) *goredis.Client {
	return goredis.NewClient(&goredis.Options{
		Addr:     c.Host,
		Username: c.User,
		Password: c.Pass,
	})
}
