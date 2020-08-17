package main

import (
	"flag"
)

// Process process dump testing all images.
// A new dump is create with a same name concatenated "-result" in end name.
func Process(dumpPath string) bool {
	if FileExists(dumpPath) == false {
		panic("Error! Argument dump doesn't informad or file not found. Execute -h to help.")
	}

	outputToValidator, outputToProductGroup := ImageAggregator(ParserProductImage(ReadFileByLine(dumpPath)))
	outputProductMap := ProductAggregator(MergeImageTypeChannel(CheckImage(outputToValidator), outputToProductGroup))
	finished := WriteFile(ParserToString(outputProductMap), dumpPath+"-result")
	result := <-finished
	return result
}

func main() {
	dumpPath := flag.String("dump", "", "Path of dump")
	flag.Parse()
	Process(*dumpPath)
}
