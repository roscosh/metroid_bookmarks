package session

import (
	"metroid_bookmarks/pkg/redispool"
	"metroid_bookmarks/pkg/session"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

const SessionKey = "session"

type Redis struct {
	redis *redispool.Redis
}

func NewRedis(pool *redis.Pool, keyPrefix string) *Redis {
	return &Redis{redis: redispool.NewRedis(pool, keyPrefix)}
}

func (r *Redis) Get(key string) (int, error) {
	userID, err := r.redis.GETEX(key, session.AnonymousExpires)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(string(userID))
}

func (r *Redis) Create(key string, value int) (bool, error) {
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

func (r *Redis) Update(key string, value, ttl int) {
	r.redis.SETEX(key, value, ttl)
}
