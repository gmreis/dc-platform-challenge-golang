package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteFile(t *testing.T) {
	fileTestWriteFile := "mock/tmpTestWriteFile"
	lines := []string{"123", "456", "789"}

	input := make(chan string)
	finished := WriteFile(input, fileTestWriteFile)
	for _, line := range lines {
		input <- line
	}
	close(input)
	<-finished

	assert.Equal(t, true, FileExists(fileTestWriteFile), "mock/tmpTestWriteFile should exist")

	lineNumber := 0
	chanReadByLine := ReadFileByLine(fileTestWriteFile)
	for line := range chanReadByLine {
		assert.Equal(t, lines[lineNumber], line, "File line should be the same as the input")
		lineNumber++
	}

	os.Remove(fileTestWriteFile)

	defer func() {
		assert.NotEqual(t, nil, recover(), "Writefile should fail")
	}()

	WriteFile(input, "")
}
