package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/rafaeljusto/redigomock"
	"github.com/stretchr/testify/assert"
)

func TestSetupRouter(t *testing.T) {

	duration := time.Duration(600) * time.Second

	getPool := func(conn redis.Conn) *redis.Pool {
		return &redis.Pool{
			Dial: func() (redis.Conn, error) { return conn, nil },
		}
	}

	t.Run("Test a request allowed", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		conn := redigomock.NewConn()

		closeCalled := false
		conn.CloseMock = func() error {
			closeCalled = true
			return nil
		}

		conn.Command("EXISTS", "ceabdbc1ee13c05bcdd8a64d2401e32755d9b0efa96c4f4f8a696cd73af37ca6").Handle(func(args []interface{}) (interface{}, error) {
			return int64(0), nil
		})

		setKeyCalled := false
		conn.Command("SET", "ceabdbc1ee13c05bcdd8a64d2401e32755d9b0efa96c4f4f8a696cd73af37ca6", true, "EX", duration.Seconds()).Handle(func(args []interface{}) (interface{}, error) {
			setKeyCalled = true
			return "OK", nil
		})

		resp := httptest.NewRecorder()
		ctx, server := gin.CreateTestContext(resp)

		server = setupRouter(getPool(conn), server)

		body := []byte("{\"id\": 123}")
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/v1/products", bytes.NewReader(body))
		ctx.Request.Header.Set("Content-Type", "application/json")

		server.ServeHTTP(resp, ctx.Request)

		assert.Equal(t, http.StatusNoContent, resp.Code, "Status code should be http.StatusNoContent")
		assert.Equal(t, "", resp.Body.String(), "Body should be empty string")
		assert.Equal(t, true, setKeyCalled, "setKey should be called")
		assert.Equal(t, true, closeCalled, "conn.Close should be called")
	})

	t.Run("Test a request not allow", func(t *testing.T) {
		conn := redigomock.NewConn()

		closeCalled := false
		conn.CloseMock = func() error {
			closeCalled = true
			return nil
		}

		conn.Command("EXISTS", "ceabdbc1ee13c05bcdd8a64d2401e32755d9b0efa96c4f4f8a696cd73af37ca6").Handle(func(args []interface{}) (interface{}, error) {
			return int64(1), nil
		})

		setKeyCalled := false
		conn.Command("SET", "ceabdbc1ee13c05bcdd8a64d2401e32755d9b0efa96c4f4f8a696cd73af37ca6", true, "EX", duration.Seconds()).Handle(func(args []interface{}) (interface{}, error) {
			setKeyCalled = true
			return "OK", nil
		})

		resp := httptest.NewRecorder()
		ctx, server := gin.CreateTestContext(resp)

		server = setupRouter(getPool(conn), server)

		body := []byte("{\"id\": 123}")
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/v1/products", bytes.NewReader(body))
		ctx.Request.Header.Set("Content-Type", "application/json")

		server.ServeHTTP(resp, ctx.Request)

		assert.Equal(t, http.StatusForbidden, resp.Code, "Status code should be http.StatusForbidden")
		assert.Equal(t, "", resp.Body.String(), "Body should be empty string")
		assert.Equal(t, false, setKeyCalled, "setKey should not be called")
		assert.Equal(t, true, closeCalled, "conn.Close should be called")
	})

	t.Run("Test a route not found", func(t *testing.T) {
		conn := redigomock.NewConn()

		closeCalled := false
		conn.CloseMock = func() error {
			closeCalled = true
			return nil
		}

		conn.Command("EXISTS", "ceabdbc1ee13c05bcdd8a64d2401e32755d9b0efa96c4f4f8a696cd73af37ca6").Handle(func(args []interface{}) (interface{}, error) {
			return int64(1), nil
		})

		setKeyCalled := false
		conn.Command("SET", "ceabdbc1ee13c05bcdd8a64d2401e32755d9b0efa96c4f4f8a696cd73af37ca6", true, "EX", duration.Seconds()).Handle(func(args []interface{}) (interface{}, error) {
			setKeyCalled = true
			return "OK", nil
		})

		resp := httptest.NewRecorder()
		ctx, server := gin.CreateTestContext(resp)

		server = setupRouter(getPool(conn), server)

		body := []byte("{\"id\": 123}")
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/v1", bytes.NewReader(body))
		ctx.Request.Header.Set("Content-Type", "application/json")

		server.ServeHTTP(resp, ctx.Request)

		assert.Equal(t, http.StatusNotFound, resp.Code, "Status code should be http.StatusNotFound")
		assert.Equal(t, "", resp.Body.String(), "Body should be empty string")
		assert.Equal(t, false, setKeyCalled, "setKey should not be called")
		assert.Equal(t, false, closeCalled, "conn.Close should not be called")
	})
}
