package main

import (
	"encoding/json"
)

// ParserToString parse ProductAggregator to string.
func ParserToString(input chan ProductMap) (output chan string) {
	output = make(chan string, 1024)

	go func() {
		for productMap := range input {
			for _, product := range productMap {
				data, _ := json.Marshal(product)
				output <- string(data)
			}
		}
		close(output)
	}()
	return

}
