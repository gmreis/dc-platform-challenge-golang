package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductType(t *testing.T) {
	expected := []string{
		"https://domain.com/image123.png",
		"https://domain.com/image456.png",
	}

	product := &ProductType{
		ProductID: "123",
		Images:    []string{"https://domain.com/image123.png"},
	}
	product.SafeAppend("https://domain.com/image456.png")

	assert.Equal(t, true, reflect.DeepEqual(expected, product.Images), "Images should have a new image")
}

func TestProductAggregator(t *testing.T) {

	newImageType := func(image string, productIDs []string, status int) *ImageType {
		return &ImageType{
			URL:        image,
			ProductIDs: productIDs,
			Status:     status,
		}
	}

	input := make(chan *ImageType, 3)
	output := ProductAggregator(input)

	input <- newImageType("https://domain.com/image111.png", []string{"123", "456"}, 200)
	input <- newImageType("https://domain.com/image222.png", []string{"456"}, 404)
	input <- newImageType("https://domain.com/image333.png", []string{"123"}, 200)

	close(input)
	var productMap ProductMap
	productMap = <-output

	expectedProduct123 := &ProductType{
		ProductID: "123",
		Images: []string{
			"https://domain.com/image111.png",
			"https://domain.com/image333.png",
		},
	}
	expectedProduct456 := &ProductType{
		ProductID: "456",
		Images: []string{
			"https://domain.com/image111.png",
		},
	}

	product123, _ := productMap["123"]
	product456, _ := productMap["456"]

	assert.Equal(t, true, reflect.DeepEqual(expectedProduct123, product123), "Product 123 should have two images")
	assert.Equal(t, true, reflect.DeepEqual(expectedProduct456, product456), "Product 456 should have one images")
	assert.Equal(t, 2, len(productMap), "ProductMap should have two items")
}
