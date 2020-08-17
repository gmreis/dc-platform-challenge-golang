package main

import (
	"fmt"
	"os"
)

// WriteFile write in file and return boolean after finished.
func WriteFile(input chan string, fileName string) (output chan bool) {
	file, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	output = make(chan bool)
	go func() {
		for data := range input {
			_, err = fmt.Fprintln(file, data)
		}
		output <- true
		close(output)
	}()

	return
}
