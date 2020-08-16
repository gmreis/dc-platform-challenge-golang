package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// WriteFile ...
func WriteFile(input chan *ProductMap, dumpName string) {
	file, err := os.Create(dumpName + "-result")

	if err != nil {
		panic(err)
	}

	productMap := <-input
	for _, product := range *productMap {
		data, _ := json.Marshal(product)
		_, err = fmt.Fprintln(file, string(data))

		if err != nil {
			panic(err)
		}
	}

}
