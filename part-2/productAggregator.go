package main

import (
	"sync"
)

// ProductType ...
type ProductType struct {
	sync.Mutex
	ProductID string
	Images    []string
}

// SafeAppend ...
func (product *ProductType) SafeAppend(image string) {
	product.Lock()
	defer product.Unlock()
	product.Images = append(product.Images, image)
}

// ProductMap ...
type ProductMap map[string]*ProductType

// ProductAggregator ...
func ProductAggregator(input chan *ImageType) (output chan *ProductMap) {
	output = make(chan *ProductMap)
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
		output <- &productMap
		close(output)
	}()

	return
}
