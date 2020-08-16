package main

import (
	"flag"
	"os"
)

func fileExists(filepath string) bool {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func process(dumpPath string) {
	if fileExists(dumpPath) {
		outputToValidator, outputToProductGroup := ImageAggregator(ParserProductImage(ReadFileByLine(dumpPath)))
		outputProductMap := ProductAggregator(MergeImageTypeChannel(CheckImage(outputToValidator), outputToProductGroup))
		WriteFile(outputProductMap, dumpPath)
	} else {
		panic("Error! Argument dump doesn't informad or file not found. Execute -h to help.")
	}
}

func main() {
	dumpPath := flag.String("dump", "", "Path of dump")
	flag.Parse()
	process(*dumpPath)
}
