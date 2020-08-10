package main

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func getDuration(str string, secondsDefault int) time.Duration {
	seconds, err := strconv.Atoi(str)
	if err != nil {
		seconds = secondsDefault
	}

	return time.Duration(seconds) * time.Second
}

func getHashBody(c *gin.Context) string {
	body := getBody(c.Request)
	sum := sha256.Sum256(body)
	return hex.EncodeToString(sum[0:])
}

func abortRequest(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}

func setupRouter(poolRedis *redis.Pool, router *gin.Engine) *gin.Engine {
	limiterMiddleware := newLimiter(
		poolRedis,
		getHashBody,
		abortRequest,
		getDuration(os.Getenv("LIMITER"), 600),
	)

	router.POST("/v1/products",
		validateMiddleware,
		limiterMiddleware,
		func(c *gin.Context) {
			c.Status(http.StatusNoContent)
		})

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	return router
}

func main() {
	redisHost := getEnv("REDIS_HOST", "localhost:6379")
	poolRedis := newRedis(redisHost)
	server := setupRouter(poolRedis, gin.Default())
	server.Run()
}
