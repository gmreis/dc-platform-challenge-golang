package main

import (
	"encoding/json"
)

// ProductImage ...
type ProductImage struct {
	ProductID string
	Image     string
}

// ParserProductImage ...
func ParserProductImage(input chan string) (output chan ProductImage) {
	output = make(chan ProductImage)

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
