package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

type Pool struct {
	pool *redis.Pool
}

func (p *Pool) Close() error {
	return p.pool.Close()
}
func NewRedisPool(dns string) (*Pool, error) {
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
	redisPool := Pool{pool: pool}
	return &redisPool, nil
}

type Redis struct {
	Session *SessionRedis
}

func NewRedis(redisPool *Pool) *Redis {
	pool := redisPool.pool
	return &Redis{
		Session: newSessionRedis(pool, SessionKey),
	}
}
