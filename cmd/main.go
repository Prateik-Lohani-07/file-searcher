package main

import (
	"fmt"
	"flag"
	"log"
	"os"
)

var Query string
var _ = fmt.Print

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		log.Fatal("Must pass in at least 2 args (query and file pattern)")
		os.Exit(1)
	}

	glob, paths := []string{}, []string{}

	Query, glob = args[0], args[1:]

	for _, g := range glob {
		expandedGlob, err := expandGlob(g)

		if err != nil {
			log.Fatalf("Error while parsing glob pattern: %v", err)
			os.Exit(1)
		}

		paths = append(paths, expandedGlob...)
	}
	
	results, err := searchAllPaths(paths)

	if err != nil {
		log.Fatalf("Failed to search for files: %v", err)
	}

	displayResults(results)
}