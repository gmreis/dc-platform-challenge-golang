package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileByLine(t *testing.T) {
	fileTestReadFile := "mock/testReadFile"
	expectedLines := []string{
		"{\"productId\":\"pid1\",\"image\":\"http://localhost:4567/images/167412.png\"}",
		"{\"productId\":\"pid2\",\"image\":\"http://localhost:4567/images/167410.png\"}",
		"{\"productId\":\"pid3\",\"image\":\"http://localhost:4567/images/167413.png\"}",
	}

	lineNumber := 0
	lines := ReadFileByLine(fileTestReadFile)
	for line := range lines {
		assert.Equal(t, expectedLines[lineNumber], line, "File line should be the same as the mock file")
		lineNumber++
	}

	defer func() {
		assert.NotEqual(t, nil, recover(), "ReadFileByLine should fail")
	}()

	ReadFileByLine("")
}
