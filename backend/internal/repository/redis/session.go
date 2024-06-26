package redis

import (
	"metroid_bookmarks/pkg/redispool"
	"metroid_bookmarks/pkg/session"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

const SessionKey = "session"

type SessionRedis struct {
	redis *redispool.Redis
}

func newSessionRedis(pool *redis.Pool, keyPrefix string) *SessionRedis {
	return &SessionRedis{redis: redispool.NewRedis(pool, keyPrefix)}
}

func (r *SessionRedis) Get(key string) (int, error) {
	userId, err := r.redis.GETEX(key, session.AnonymousExpires)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(userId))
}

func (r *SessionRedis) Create(key string, value int) (bool, error) {
	script := `
	if redis.call('exists', KEYS[1]) == 0 then
	redis.call('setex', KEYS[1], 3600, ARGV[1])
	return 1
	else
	return 0
	end
	`
	return r.redis.EVAL(key, script, 1, value)
}

func (r *SessionRedis) Update(key string, value, ttl int) {
	r.redis.SETEX(key, value, ttl)
}
