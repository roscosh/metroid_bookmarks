package redis

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"metroid_bookmarks/pkg/misc"
	"time"
)

var logger = misc.GetLogger()

type RedisPool struct {
	pool *redis.Pool
	ctx  context.Context
}

func (d *RedisPool) Close() error {
	return d.pool.Close()
}
func NewRedisPool(dns string) (*RedisPool, error) {
	pool := &redis.Pool{
		MaxIdle:     10,
		MaxActive:   0,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(dns)
			if err != nil {
				panic(err)
			}
			return c, err
		},
	}

	conn := pool.Get()
	defer conn.Close()
	if _, err := conn.Do("ping"); err != nil {
		panic(err)
	}
	redisPool := RedisPool{pool: pool, ctx: context.Background()}
	return &redisPool, nil
}

type Redis struct {
	Session *SessionRedis
}

func NewRedis(redisPool *RedisPool) *Redis {
	pool := redisPool.pool
	return &Redis{
		Session: newSessionRedis(pool, SessionKey),
	}
}
