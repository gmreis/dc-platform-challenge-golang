package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// ReadFileByLine read file and return line by line.
func ReadFileByLine(filePath string) (output chan string) {
	output = make(chan string, 1024)

	file, err := os.Open(filePath)

	if err != nil {
		panic("Failed to open the file")
	} else {
		go func() {
			reader := bufio.NewReader(file)
			var line string
			for {
				line, err = reader.ReadString('\n')

				line = strings.TrimSpace(line)
				if len(line) > 0 {
					output <- line
				}

				if err != nil {
					break
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
