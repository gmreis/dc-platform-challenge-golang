package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStatusByRequest(t *testing.T) {

	getHanderStatus := func(statusCode int) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(statusCode)
		})
	}

	tsOK := httptest.NewServer(getHanderStatus(http.StatusOK))
	tsNotFound := httptest.NewServer(getHanderStatus(http.StatusNotFound))

	defer tsOK.Close()
	defer tsNotFound.Close()

	input := make(chan *ImageType)

	chanCheckImage := CheckImage(input)

	expectedStatus := make(map[string]int)
	expectedStatus[tsOK.URL] = http.StatusOK
	expectedStatus[tsNotFound.URL] = http.StatusNotFound
	expectedStatus[""] = 0

	for url := range expectedStatus {
		input <- &ImageType{URL: url}
	}
	close(input)

	for image := range chanCheckImage {
		assert.Equal(t, expectedStatus[image.URL], image.Status, "Status should be "+string(expectedStatus[image.URL]))
	}
}
