package main

import (
	"io/fs"
	"os"
	"path/filepath"
)


func searchAllPaths(paths []string) ([]*SearchResult, error) {
	var allResults = []*SearchResult{}
	
	for _, path := range paths {
		pathResults, err := searchPath(path)
		if err != nil {
			return nil, err
		}

		allResults = append(allResults, pathResults...)
	}

	return allResults, nil
}

func searchPath(path string) ([]*SearchResult, error) {
	var results = []*SearchResult{}

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
		
		if lin, col, content, err := searchQueryFound(path); err == nil {
			results = append(results, &SearchResult{lin, col, path, content})
		}

		return nil
	}

	if *RecursiveSearch {
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

