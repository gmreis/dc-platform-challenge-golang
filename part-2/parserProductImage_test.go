package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserProductImage(t *testing.T) {
	fileTestReadFile := "mock/testReadFile"
	expectedProductImage := []ProductImage{
		{ProductID: "pid1", Image: "http://localhost:4567/images/167412.png"},
	}

	chanProductImage := ParserProductImage(ReadFileByLine(fileTestReadFile))
	productImage := <-chanProductImage
	assert.Equal(t, true, reflect.DeepEqual(expectedProductImage[0], productImage), "ProductImage should be the same as the mock")
}
