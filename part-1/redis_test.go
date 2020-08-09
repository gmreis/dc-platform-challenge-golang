package main

import (
	"testing"
	"time"

	"github.com/rafaeljusto/redigomock"
	"github.com/stretchr/testify/assert"
)

func TestExistKeyRedis(t *testing.T) {

	t.Run("Return TRUE if key exists", func(t *testing.T) {
		conn := redigomock.NewConn()
		conn.Command("EXISTS", "123").Handle(func(args []interface{}) (interface{}, error) {
			return int64(1), nil
		})

		ok, err := existKey(conn, "123")

		assert.Equal(t, true, ok, "ok should be true")
		assert.Equal(t, nil, err, "err should be nil")
	})

	t.Run("Return FALSE if key doesn't exist", func(t *testing.T) {
		conn := redigomock.NewConn()
		conn.Command("EXISTS", "123").Handle(func(args []interface{}) (interface{}, error) {
			return int64(0), nil
		})

		ok, err := existKey(conn, "123")

		assert.Equal(t, false, ok, "ok should be true")
		assert.Equal(t, nil, err, "err should be nil")
	})
}

func TestSetKeyRedis(t *testing.T) {
	conn := redigomock.NewConn()
	conn.Command("SET", "123", true, "EX", time.Minute.Seconds()).Handle(func(args []interface{}) (interface{}, error) {
		return "OK", nil
	})

	err := setKey(conn, "123", time.Minute)

	assert.Equal(t, nil, err, "err should be nil")
}
