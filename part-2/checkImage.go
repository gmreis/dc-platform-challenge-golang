package main

import (
	"net/http"
)

func getStatusByRequest(url string) int {
	resp, err := http.Get(url)

	if err != nil {
		return 0
	}

	defer resp.Body.Close()
	return resp.StatusCode
}

// CheckImage ...
func CheckImage(input chan *ImageType) (output chan *ImageType) {
	output = make(chan *ImageType)

	go func() {
		for image := range input {
			image.SafeSetStatus(getStatusByRequest(image.URL))
			output <- image
		}
		close(output)
	}()

	return
}
