package main

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

func newRedis(server string) *redis.Pool {
	var pool = redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return &pool
}

func existKey(conn redis.Conn, key string) (bool, error) {
	exist, err := conn.Do("EXISTS", key)
	return redis.Bool(exist, err)
}

func setKey(conn redis.Conn, key string, limiter time.Duration) error {
	_, err := redis.String(conn.Do("SET", key, true, "EX", limiter.Seconds()))
	return err
}
