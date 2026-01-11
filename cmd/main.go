package main

import (
	"log"
	"os"
)

var Query string

func main() {
	q, paths, err := parseArgs()
	
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	Query = q
	
	results, err := searchAllPaths(paths)
	if err != nil {
		log.Fatalf("Failed to search for files: %v", err)
	}

	displayResults(results)
}
