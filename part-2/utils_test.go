package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {
	assert.Equal(t, true, FileExists("mock/testReadFile"), "mock/testReadFile should exist")
	assert.Equal(t, false, FileExists("NotExistFile"), "NotExistFile should not exist")
}
