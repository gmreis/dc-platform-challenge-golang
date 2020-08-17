package main

import (
	"sync"
)

// ImageType with status and productIds that use this image.
type ImageType struct {
	sync.Mutex
	URL        string
	ProductIDs []string
	Status     int
}

// SafeAppendProductID append productId in image.
func (image *ImageType) SafeAppendProductID(productID string) {
	image.Lock()
	defer image.Unlock()
	image.ProductIDs = append(image.ProductIDs, productID)
}

// SafeGetAndRemovedProductIDs return productId list and remove products in image.
func (image *ImageType) SafeGetAndRemovedProductIDs() (productIDs []string) {
	image.Lock()
	defer image.Unlock()
	productIDs = image.ProductIDs
	image.ProductIDs = []string{}
	return
}

// SafeSetStatus set status code of image.
func (image *ImageType) SafeSetStatus(status int) {
	image.Lock()
	defer image.Unlock()
	image.Status = status
}

// SafeGetStatus return status code of image.
func (image *ImageType) SafeGetStatus() (status int) {
	image.Lock()
	defer image.Unlock()
	status = image.Status
	return
}

// ImageAggregator aggregate productId of same image in a map.
// If a image is a first time, send it to validator url.
// Case a image exists, send to product aggregator.
func ImageAggregator(input chan ProductImage) (outputValidator chan *ImageType, outputProductAggregator chan *ImageType) {
	outputValidator = make(chan *ImageType, 1024)
	outputProductAggregator = make(chan *ImageType, 1024)

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
					outputProductAggregator <- image
				}
			}
		}

		close(outputValidator)
		close(outputProductAggregator)
	}()

	return
}
