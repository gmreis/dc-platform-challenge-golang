package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// ReadFileByLine ...
func ReadFileByLine(filePath string) (output chan string) {
	output = make(chan string)

	file, err := os.Open(filePath)

	if err != nil {
		panic("Failed to open the file")
	} else {
		go func() {
			reader := bufio.NewReader(file)
			var line string
			for {
				line, err = reader.ReadString('\n')

				if err != nil {
					break
				}

				line = strings.TrimSpace(line)

				if len(line) > 0 {
					output <- line
				}
			}

			if err != io.EOF {
				panic("Failed to read the file")
			}

			file.Close()
			close(output)
		}()
	}

	return
}
