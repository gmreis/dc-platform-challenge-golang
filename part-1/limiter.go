package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

// KeyDefine is a function used to defined the key of cache.
type KeyDefine func(*gin.Context) string

// AbortCallback is a function used to abort the request.
type AbortCallback func(*gin.Context)

func newLimiter(pool *redis.Pool, key KeyDefine, abort AbortCallback, limiter time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn := pool.Get()
		defer conn.Close()

		k := key(c)
		ok, err := existKey(conn, k)
		if err != nil {
			panic(err)
		}

		if ok == false {
			c.Next()
			setKey(conn, k, limiter)
		} else {
			abort(c)
		}
	}
}
