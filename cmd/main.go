package main

import (
	"log"
	"os"

	"filesearch/internal/search"
	"filesearch/internal/cli"
)

func main() {
	query, paths, err := cli.ParseArgs()
	
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	results, err := search.SearchAllPaths(query, paths)

	if err != nil {
		log.Fatalf("Failed to search for files: %v", err)
		os.Exit(1)
	}

	cli.DisplayResults(query, results)
}
