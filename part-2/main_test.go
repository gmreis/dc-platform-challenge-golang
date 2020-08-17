package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	result := Process("mock/testReadFile")

	assert.Equal(t, true, result, "Result should be true")
	assert.Equal(t, true, FileExists("mock/testReadFile-result"), "testReadFile-result should exist")

	os.Remove("mock/testReadFile-result")

	defer func() {
		assert.NotEqual(t, nil, recover(), "Process should fail")
	}()

	Process("")
}
