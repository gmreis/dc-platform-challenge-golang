package main

import (
	"sync"
)

// ImageType ...
type ImageType struct {
	sync.Mutex
	URL        string
	ProductIDs []string
	Status     int
}

// SafeAppendProductID ...
func (image *ImageType) SafeAppendProductID(productID string) {
	image.Lock()
	defer image.Unlock()
	image.ProductIDs = append(image.ProductIDs, productID)
}

// SafeGetAndRemovedProductIDs ...
func (image *ImageType) SafeGetAndRemovedProductIDs() (productIDs []string) {
	image.Lock()
	defer image.Unlock()
	productIDs = image.ProductIDs
	image.ProductIDs = []string{}
	return
}

// SafeSetStatus ...
func (image *ImageType) SafeSetStatus(status int) {
	image.Lock()
	defer image.Unlock()
	image.Status = status
}

// SafeGetStatus ...
func (image *ImageType) SafeGetStatus() int {
	image.Lock()
	defer image.Unlock()
	return image.Status
}

// ImageAggregator ...
func ImageAggregator(input chan ProductImage) (outputValidator chan *ImageType, outputProductGroup chan *ImageType) {
	outputValidator = make(chan *ImageType)
	outputProductGroup = make(chan *ImageType)

	imageMap := make(map[string]*ImageType)

	go func() {
		for productImage := range input {
			image, hasImage := imageMap[productImage.Image]

			if hasImage == false {
				image = &ImageType{
					URL:        productImage.Image,
					ProductIDs: []string{productImage.ProductID},
				}

				imageMap[productImage.Image] = image

				outputValidator <- image
			} else {
				image.SafeAppendProductID(productImage.ProductID)
				if image.SafeGetStatus() > 0 {
					outputProductGroup <- image
				}
			}
		}

		close(outputValidator)
		close(outputProductGroup)
	}()

	return
}
