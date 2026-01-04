package main

import (
	// "fmt"
	"flag"
	"log"
)


func main() {
	flag.Parse()
	searchResults, err := searchDir()

	if err != nil {
		log.Fatalf("Failed to search for files: %v", err)
	}

	displayResults(searchResults)
}