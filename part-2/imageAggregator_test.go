package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImageType(t *testing.T) {

	image := &ImageType{
		URL:        "https://domain.com/image.png",
		ProductIDs: []string{"123"},
	}

	productIDs := image.SafeGetAndRemovedProductIDs()
	assert.Equal(t, true, reflect.DeepEqual([]string{"123"}, productIDs), "ProductIDs should be 123")
	assert.Equal(t, true, reflect.DeepEqual([]string{}, image.ProductIDs), "ProductIDs should be empty")
	image.SafeAppendProductID("789")
	assert.Equal(t, true, reflect.DeepEqual([]string{"789"}, image.ProductIDs), "ProductIDs should be 789")
	image.SafeSetStatus(404)
	assert.Equal(t, 404, image.SafeGetStatus(), "Status should be 404")
}

func TestImageAggregator(t *testing.T) {

	newProductImage := func(productID string, image string) ProductImage {
		return ProductImage{
			ProductID: productID,
			Image:     image,
		}
	}

	input := make(chan ProductImage, 4)
	outputValidator, outputProductGroup := ImageAggregator(input)

	var image *ImageType
	input <- newProductImage("123", "https://domain.com/image111.png")
	image = <-outputValidator
	image.SafeSetStatus(200)
	assert.Equal(t, "https://domain.com/image111.png", image.URL, "Image should be returned in outputValidator")

	input <- newProductImage("456", "https://domain.com/image444.png")
	input <- newProductImage("123", "https://domain.com/image444.png")

	image = <-outputValidator
	assert.Equal(t, "https://domain.com/image444.png", image.URL, "Image should be returned in outputValidator")

	input <- newProductImage("456", "https://domain.com/image111.png")
	image = <-outputProductGroup

	assert.Equal(t, "https://domain.com/image111.png", image.URL, "Image should be returned in outputProductGroup")
	assert.Equal(t, true, reflect.DeepEqual([]string{"123", "456"}, image.ProductIDs), "Image should have many ids")

	close(input)
	assert.Equal(t, 0, len(outputValidator), "OutputValidator should be empty")
	assert.Equal(t, 0, len(outputProductGroup), "OutputValidator should be empty")
}
