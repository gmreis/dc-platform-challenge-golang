package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeImageTypeChannel(t *testing.T) {
	input1 := make(chan *ImageType)
	input2 := make(chan *ImageType)
	input3 := make(chan *ImageType)

	imageMap := make(map[string]*ImageType)
	imageMap["https://domain.com/image111.png"] = &ImageType{
		URL:        "https://domain.com/image111.png",
		ProductIDs: []string{"111"},
	}
	imageMap["https://domain.com/image222.png"] = &ImageType{
		URL:        "https://domain.com/image222.png",
		ProductIDs: []string{"222"},
	}
	imageMap["https://domain.com/image333.png"] = &ImageType{
		URL:        "https://domain.com/image333.png",
		ProductIDs: []string{"333"},
	}

	output := MergeImageTypeChannel(input1, input2, input3)

	input1 <- imageMap["https://domain.com/image111.png"]
	input2 <- imageMap["https://domain.com/image222.png"]
	input3 <- imageMap["https://domain.com/image333.png"]

	close(input1)
	close(input2)
	close(input3)

	for image := range output {
		assert.Equal(t, true, reflect.DeepEqual(imageMap[image.URL], image), "All inputs should be returned in output")
	}
}
