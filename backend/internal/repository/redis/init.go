package redis

import (
	"github.com/gomodule/redigo/redis"
	"metroid_bookmarks/pkg/misc"
	"time"
)

var logger = misc.GetLogger()

func NewRedisPool(dns string) (*redis.Pool, error) {
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

	return pool, nil
}

type Redis struct {
	Session *SessionRedis
}

func NewRedis(pool *redis.Pool) *Redis {
	return &Redis{
		Session: newSessionRedis(pool, SessionKey),
	}
}
