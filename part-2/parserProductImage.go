package main

import (
	"encoding/json"
)

// ProductImage is a struct that it's received by dump.
type ProductImage struct {
	ProductID string
	Image     string
}

// ParserProductImage parse string to ProductImage.
func ParserProductImage(input chan string) (output chan ProductImage) {
	output = make(chan ProductImage, 1024)

	go func() {
		for productImageJSON := range input {
			var productImage ProductImage
			json.Unmarshal([]byte(productImageJSON), &productImage)
			output <- productImage
		}
		close(output)
	}()

	return
}
