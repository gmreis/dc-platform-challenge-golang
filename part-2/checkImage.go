package main

import (
	"net/http"
)

// GetStatusByRequest make request in URL and return status code.
func GetStatusByRequest(url string) int {
	resp, err := http.Get(url)

	if err != nil {
		return 0
	}

	defer resp.Body.Close()
	return resp.StatusCode
}

// CheckImage check status of URL.
func CheckImage(input chan *ImageType) (output chan *ImageType) {
	output = make(chan *ImageType, 1024)

	go func() {
		for image := range input {
			image.SafeSetStatus(GetStatusByRequest(image.URL))
			output <- image
		}
		close(output)
	}()

	return
}
