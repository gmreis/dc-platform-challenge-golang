package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserToString(t *testing.T) {
	mockProductMap := make(ProductMap)
	mockProductMap["pid1"] = &ProductType{
		ProductID: "pid1",
		Images:    []string{"http://localhost:4567/images/167412.png"},
	}

	expectedProductString := "{\"productId\":\"pid1\",\"images\":[\"http://localhost:4567/images/167412.png\"]}"

	input := make(chan ProductMap)
	chanProductString := ParserToString(input)
	input <- mockProductMap
	close(input)

	assert.Equal(t, expectedProductString, <-chanProductString, "ProductString should be equal a expectedProductString")
}
