package search

import (
	"io/fs"
	"os"
	"path/filepath"

	"filesearch/internal/cli"
)

var Query string

func SearchAllPaths(query string, paths []string) ([]SearchResult, error) {
	var allResults = []SearchResult{}
	Query = query
	
	for _, path := range paths {
		pathResults, err := searchPath(path)
		if err != nil {
			return nil, err
		}

		allResults = append(allResults, pathResults...)
	}

	return allResults, nil
}

func searchPath(path string) ([]SearchResult, error) {
	var results = []SearchResult{}

	processFileFn := func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			if info != nil && info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		ignore := toIgnore(info.Name())
		if info.Mode().IsDir() || ignore || !info.Mode().IsRegular() {
			if ignore {
				return filepath.SkipDir
			}

			return nil
		}
		
		r, err := searchQueryFound(path)
		if err == nil {
			results = append(results, r...)
		}

		return nil
	}

	if *cli.RecursiveSearch {
		if err := filepath.Walk(path, processFileFn); err != nil {
			return nil, err
		}
	} else {
		info, err := os.Stat(path)
		err = processFileFn(path, info, err)

		if err != nil {
			return nil, err
		}
	}

	return results, nil
}

