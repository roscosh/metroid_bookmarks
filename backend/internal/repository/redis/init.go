package redis

import (
	"metroid_bookmarks/internal/repository/redis/session"
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	MaxIdle     = 10
	MaxActive   = 0
	IdleTimeout = 240 * time.Second
)

type Pool struct {
	pool *redis.Pool
}

func (p *Pool) Close() error {
	return p.pool.Close()
}

func NewRedisPool(dns string) (*Pool, error) {
	pool := &redis.Pool{
		MaxIdle:     MaxIdle,
		MaxActive:   MaxActive,
		IdleTimeout: IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(dns)
			if err != nil {
				panic(err)
			}

			return c, nil
		},
	}

	conn := pool.Get()
	defer conn.Close()

	if _, err := conn.Do("ping"); err != nil {
		panic(err)
	}

	redisPool := Pool{pool: pool}

	return &redisPool, nil
}

type Redis struct {
	Session *session.Redis
}

func NewRedis(redisPool *Pool) *Redis {
	pool := redisPool.pool

	return &Redis{
		Session: session.NewRedis(pool, session.SessionKey),
	}
}
