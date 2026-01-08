package main

import (
	"path/filepath"
	"flag"
	"log"
	"os"
)

func parseArgs() (string, []string) {
	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		log.Fatal("Must pass in at least 2 args (query and file pattern)")
		os.Exit(1)
	}

	query, glob := args[0], args[1:]
	paths := []string{}

	for _, g := range glob {
		expandedGlob, err := expandGlob(g)

		if err != nil {
			log.Fatalf("Error while parsing glob pattern: %v", err)
			os.Exit(1)
		}

		paths = append(paths, expandedGlob...)
	}

	return query, paths
}

func expandGlob(glob string) ([]string, error) {
	matches, err := filepath.Glob(glob)
	if err != nil {
		return nil, err
	}

	return matches, nil
}