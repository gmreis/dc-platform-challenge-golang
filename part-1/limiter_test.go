package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/rafaeljusto/redigomock"
	"github.com/stretchr/testify/assert"
)

func TestNewLimiter(t *testing.T) {

	getPool := func(conn redis.Conn) *redis.Pool {
		return &redis.Pool{
			Dial: func() (redis.Conn, error) { return conn, nil },
		}
	}

	mockGetKey := func(mock string) KeyDefine {
		return func(c *gin.Context) string {
			return mock
		}
	}

	abortHandler := func(c *gin.Context) {
		c.AbortWithStatus(http.StatusBadGateway)
	}

	t.Run("newLimiter return a gin.HandlerFunc", func(t *testing.T) {
		conn := redigomock.NewConn()

		closeCalled := false
		conn.CloseMock = func() error {
			closeCalled = true
			return nil
		}

		conn.Command("EXISTS", "hash").Handle(func(args []interface{}) (interface{}, error) {
			return int64(0), nil
		})

		setKeyCalled := false
		conn.Command("SET", "hash", true, "EX", time.Minute.Seconds()).Handle(func(args []interface{}) (interface{}, error) {
			setKeyCalled = true
			return "OK", nil
		})

		pool := getPool(conn)
		limiter := newLimiter(pool, mockGetKey("hash"), abortHandler, time.Minute)

		gin.SetMode(gin.TestMode)

		resp := httptest.NewRecorder()
		ctx, server := gin.CreateTestContext(resp)

		server.Use(
			limiter,
			func(c *gin.Context) {
				c.String(http.StatusOK, "Finish")
			},
		)

		ctx.Request, _ = http.NewRequest(http.MethodPost, "/", nil)
		ctx.Request.Header.Set("Content-Type", "application/json")

		server.ServeHTTP(resp, ctx.Request)

		assert.Equal(t, http.StatusOK, resp.Code, "Status code should be 200")
		assert.Equal(t, "Finish", resp.Body.String(), "Body should be 'Finish' string")
		assert.Equal(t, true, setKeyCalled, "setKey should be called")
		assert.Equal(t, true, closeCalled, "conn.Close should be called")
	})

	t.Run("newLimiter return a gin.HandlerFunc", func(t *testing.T) {
		conn := redigomock.NewConn()

		closeCalled := false
		conn.CloseMock = func() error {
			closeCalled = true
			return nil
		}

		conn.Command("EXISTS", "hash").Handle(func(args []interface{}) (interface{}, error) {
			return int64(1), nil
		})

		setKeyCalled := false
		conn.Command("SET", "hash", true, "EX", time.Minute.Seconds()).Handle(func(args []interface{}) (interface{}, error) {
			setKeyCalled = true
			return "OK", nil
		})

		pool := getPool(conn)
		limiter := newLimiter(pool, mockGetKey("hash"), abortHandler, time.Minute)

		gin.SetMode(gin.TestMode)

		resp := httptest.NewRecorder()
		ctx, server := gin.CreateTestContext(resp)

		server.Use(limiter)

		ctx.Request, _ = http.NewRequest(http.MethodPost, "/", nil)
		ctx.Request.Header.Set("Content-Type", "application/json")

		server.ServeHTTP(resp, ctx.Request)

		assert.Equal(t, http.StatusBadGateway, resp.Code, "Status code should be 200")
		assert.Equal(t, "", resp.Body.String(), "Body should be empty string")
		assert.Equal(t, false, setKeyCalled, "setKey should be called")
		assert.Equal(t, true, closeCalled, "conn.Close should be called")
	})
}
