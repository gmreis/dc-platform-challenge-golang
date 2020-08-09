package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gmreis/go-limiter"
	"github.com/gmreis/go-limiter/drivers"
)

type applicationJSON struct {
	ContentType string `header:"Content-type" binding:"required,eq=application/json"`
}

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

func headerValidate(c *gin.Context) {
	if err := c.ShouldBindHeader(&applicationJSON{}); err != nil {
		c.AbortWithStatus(http.StatusNotAcceptable)
	}
}

func getHashBody(c *gin.Context) string {
	body, _ := c.GetRawData()
	sum := sha256.Sum256(body)
	return hex.EncodeToString(sum[0:])
}

func abortRequest(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}

func main() {
	router := gin.Default()

	redisHost := getEnv("REDIS_HOST", "localhost:6379")
	log.Println("redisHost", redisHost)
	log.Println("LIMITER", os.Getenv("LIMITER"))
	durationLimiter := getDuration(os.Getenv("LIMITER"), 600)

	limiterMiddleware := limiter.NewLimiter(
		drivers.NewRedis(redisHost),
		getHashBody,
		abortRequest,
		durationLimiter,
	)

	router.POST("/v1/products",
		headerValidate,
		limiterMiddleware,
		func(c *gin.Context) {
			c.Status(http.StatusNoContent)
		})

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	router.Run(":8080")
}
