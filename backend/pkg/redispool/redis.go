package redispool

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	pool      *redis.Pool
	keyPrefix string
}

func NewRedis(pool *redis.Pool, keyPrefix string) *Redis {
	return &Redis{pool: pool, keyPrefix: keyPrefix}
}

func (r *Redis) GET(key string) ([]byte, error) {
	conn := r.pool.Get()
	defer conn.Close()

	key = r.addPrefixKey(key, r.keyPrefix)

	return redis.Bytes(conn.Do("GET", key))
}

func (r *Redis) GETEX(key string, ttl int) ([]byte, error) {
	conn := r.pool.Get()
	defer conn.Close()

	key = r.addPrefixKey(key, r.keyPrefix)

	return redis.Bytes(conn.Do("GETEX", key, "EX", ttl))
}

func (r *Redis) SETNX(key string, value interface{}) (bool, error) {
	conn := r.pool.Get()
	defer conn.Close()

	key = r.addPrefixKey(key, r.keyPrefix)

	return redis.Bool(conn.Do("SETNX", key, value))
}

func (r *Redis) EXPIRE(key string, ttl int) {
	conn := r.pool.Get()
	defer conn.Close()

	key = r.addPrefixKey(key, r.keyPrefix)

	_, err := conn.Do("EXPIRE", key, ttl)
	if err != nil {
		panic("Ошибка подключения к Redis")
	}
}

func (r *Redis) SETEX(key string, value interface{}, ttl int) {
	conn := r.pool.Get()
	defer conn.Close()

	key = r.addPrefixKey(key, r.keyPrefix)

	_, err := conn.Do("SETEX", key, ttl, value)
	if err != nil {
		panic("Ошибка подключения к Redis")
	}
}

func (r *Redis) EVAL(key, script string, numKeys int, args ...any) (bool, error) {
	conn := r.pool.Get()
	defer conn.Close()

	key = r.addPrefixKey(key, r.keyPrefix)

	return redis.Bool(conn.Do("EVAL", script, numKeys, key, args))
}

func (r *Redis) addPrefixKey(key, prefix string) string {
	return fmt.Sprintf("%s_%s", prefix, key)
}
