package main

import (
	"fmt"
	"log"
)

var Query string
var _ = fmt.Print

func main() {
	var paths []string
	Query, paths = parseArgs()
	
	results, err := searchAllPaths(paths)
	if err != nil {
		log.Fatalf("Failed to search for files: %v", err)
	}

	displayResults(results)
}
