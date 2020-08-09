package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetBody(t *testing.T) {
	gin.SetMode(gin.TestMode)

	resp := httptest.NewRecorder()
	ctx, server := gin.CreateTestContext(resp)

	var firstBodyHandler []byte
	var secondBodyHandler []byte
	server.Use(
		func(c *gin.Context) {
			firstBodyHandler = getBody(c.Request)
			c.Next()
		},
		func(c *gin.Context) {
			secondBodyHandler = getBody(c.Request)
			c.Status(http.StatusOK)
		},
	)

	body := []byte("abc")
	ctx.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewReader(body))

	server.ServeHTTP(resp, ctx.Request)

	assert.Equal(t, "abc", string(firstBodyHandler), "firstBodyHandler should be equal 'abc'")
	assert.Equal(t, "abc", string(secondBodyHandler), "secondBodyHandler should be equal 'abc'")
	assert.Equal(t, string(firstBodyHandler), string(secondBodyHandler), "firstBodyHandler should be equal secondBodyHandler")
}
