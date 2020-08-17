package main

import (
	"sync"
)

// ProductType is a product with its images.
type ProductType struct {
	sync.Mutex
	ProductID string   `json:"productId"`
	Images    []string `json:"images"`
}

// SafeAppend append a image in product.
func (product *ProductType) SafeAppend(image string) {
	product.Lock()
	defer product.Unlock()
	product.Images = append(product.Images, image)
}

// ProductMap is a map of ProductType
type ProductMap map[string]*ProductType

// ProductAggregator joint images in a product if image status is 200.
func ProductAggregator(input chan *ImageType) (output chan ProductMap) {
	output = make(chan ProductMap)
	productMap := make(ProductMap)

	go func() {
		for image := range input {
			for _, productID := range image.SafeGetAndRemovedProductIDs() {
				product, hasProduct := productMap[productID]

				if hasProduct == false {
					product = &ProductType{
						ProductID: productID,
					}
					productMap[productID] = product
				}

				if image.Status == 200 && len(product.Images) < 3 {
					product.SafeAppend(image.URL)
				}
			}
		}
		output <- productMap
		close(output)
	}()

	return
}
