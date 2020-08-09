package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func getBody(request *http.Request) []byte {
	var bodyBytes []byte
	if request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(request.Body)
		request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	return bodyBytes
}
