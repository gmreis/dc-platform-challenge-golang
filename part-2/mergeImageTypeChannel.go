package main

import (
	"sync"
)

// MergeImageTypeChannel ...
func MergeImageTypeChannel(cs ...chan *ImageType) (output chan *ImageType) {
	output = make(chan *ImageType)

	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(c <-chan *ImageType) {
			for v := range c {
				output <- v
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return
}
