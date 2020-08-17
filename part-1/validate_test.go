package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestValidateMiddleware(t *testing.T) {

	t.Run("Call next handler if body is a object JSON and Content-type is application/JSON", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		resp := httptest.NewRecorder()
		ctx, server := gin.CreateTestContext(resp)

		server.Use(
			validateMiddleware,
			func(c *gin.Context) {
				c.String(http.StatusOK, "Finish")
			},
		)

		body := []byte("{\"id\": 123}")
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
		ctx.Request.Header.Set("Content-Type", "application/json")

		server.ServeHTTP(resp, ctx.Request)

		assert.Equal(t, resp.Code, http.StatusOK, "Status code should be 200")
		assert.Equal(t, resp.Body.String(), "Finish", "Body should be 'Finish' string")
	})

	t.Run("Call next handler if body is a array JSON and Content-type is application/JSON", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		resp := httptest.NewRecorder()
		ctx, server := gin.CreateTestContext(resp)

		server.Use(
			validateMiddleware,
			func(c *gin.Context) {
				c.String(http.StatusOK, "Finish")
			},
		)

		body := []byte("[{\"id\": 123}]")
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
		ctx.Request.Header.Set("Content-Type", "application/json")

		server.ServeHTTP(resp, ctx.Request)

		assert.Equal(t, resp.Code, http.StatusOK, "Status code should be 200")
		assert.Equal(t, resp.Body.String(), "Finish", "Body should be 'Finish' string")
	})

	t.Run("Response not acceptable status if Content-type isn't application/JSON", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		resp := httptest.NewRecorder()
		ctx, server := gin.CreateTestContext(resp)

		server.Use(validateMiddleware)

		body := []byte("{\"id\": 123}")
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
		ctx.Request.Header.Set("Content-Type", "text/plain")

		server.ServeHTTP(resp, ctx.Request)

		assert.Equal(t, resp.Code, http.StatusNotAcceptable, "Status code should be 406")
		assert.Equal(t, resp.Body.String(), "", "Body should be empty string")
	})

	t.Run("Response not acceptable status if Content-type doesn't exist", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		resp := httptest.NewRecorder()
		ctx, server := gin.CreateTestContext(resp)

		server.Use(
			validateMiddleware,
			func(c *gin.Context) {
				c.Status(http.StatusOK)
			},
		)

		body := []byte("{\"id\": 123}")
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewReader(body))

		server.ServeHTTP(resp, ctx.Request)

		assert.Equal(t, resp.Code, http.StatusNotAcceptable, "Status code should be 406")
		assert.Equal(t, resp.Body.String(), "", "Body should be empty string")
	})

	t.Run("Response not acceptable status if body is empty", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		resp := httptest.NewRecorder()
		ctx, server := gin.CreateTestContext(resp)

		server.Use(validateMiddleware)

		body := []byte("")
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
		ctx.Request.Header.Set("Content-Type", "application/json")

		server.ServeHTTP(resp, ctx.Request)

		assert.Equal(t, resp.Code, http.StatusNotAcceptable, "Status code should be 406")
		assert.Equal(t, resp.Body.String(), "", "Body should be empty string")
	})

	t.Run("Response not acceptable status if body isn't JSON object", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		resp := httptest.NewRecorder()
		ctx, server := gin.CreateTestContext(resp)

		server.Use(validateMiddleware)

		body := []byte("123")
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
		ctx.Request.Header.Set("Content-Type", "application/json")

		server.ServeHTTP(resp, ctx.Request)

		assert.Equal(t, resp.Code, http.StatusNotAcceptable, "Status code should be 406")
		assert.Equal(t, resp.Body.String(), "", "Body should be empty string")
	})

	t.Run("Response not acceptable status if body is a empty JSON object", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		resp := httptest.NewRecorder()
		ctx, server := gin.CreateTestContext(resp)

		server.Use(validateMiddleware)

		body := []byte("{}")
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
		ctx.Request.Header.Set("Content-Type", "application/json")

		server.ServeHTTP(resp, ctx.Request)

		assert.Equal(t, resp.Code, http.StatusNotAcceptable, "Status code should be 406")
		assert.Equal(t, resp.Body.String(), "", "Body should be empty string")
	})
}
