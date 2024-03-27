package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type baseRedis struct {
	pool      *redis.Pool
	keyPrefix string
}

func newBaseRedis(pool *redis.Pool, keyPrefix string) *baseRedis {
	return &baseRedis{pool: pool, keyPrefix: keyPrefix}
}

func (r *baseRedis) GET(key string) ([]byte, error) {
	conn := r.pool.Get()
	defer conn.Close()
	key = r.addPrefixKey(key, r.keyPrefix)
	return redis.Bytes(conn.Do("GET", key))
}

func (r *baseRedis) GETEX(key string, TTL int) ([]byte, error) {
	conn := r.pool.Get()
	defer conn.Close()
	key = r.addPrefixKey(key, r.keyPrefix)
	return redis.Bytes(conn.Do("GETEX", key, "EX", TTL))
}

func (r *baseRedis) SETNX(key string, value interface{}) (bool, error) {
	conn := r.pool.Get()
	defer conn.Close()
	key = r.addPrefixKey(key, r.keyPrefix)
	return redis.Bool(conn.Do("SETNX", key, value))
}

func (r *baseRedis) EXPIRE(key string, TTL int) {
	conn := r.pool.Get()
	defer conn.Close()
	key = r.addPrefixKey(key, r.keyPrefix)
	_, err := conn.Do("EXPIRE", key, TTL)
	if err != nil {
		panic("Ошибка подключения к Redis")
	}
}

func (r *baseRedis) SETEX(key string, value interface{}, TTL int) {
	conn := r.pool.Get()
	defer conn.Close()
	key = r.addPrefixKey(key, r.keyPrefix)
	_, err := conn.Do("SETEX", key, TTL, value)
	if err != nil {
		panic("Ошибка подключения к Redis")
	}
}

func (r *baseRedis) EVAL(key string, script string, numKeys int, args ...any) (bool, error) {
	conn := r.pool.Get()
	defer conn.Close()
	key = r.addPrefixKey(key, r.keyPrefix)
	return redis.Bool(conn.Do("EVAL", script, numKeys, key, args))
}

func (r *baseRedis) Keys() {
	conn := r.pool.Get()
	defer conn.Close()
	keys, _ := redis.Strings(conn.Do("KEYS", "*"))
	for _, key := range keys {
		logger.Info(key)
	}
}

func (r *baseRedis) addPrefixKey(key, prefix string) string {
	return fmt.Sprintf("%s_%s", prefix, key)
}
