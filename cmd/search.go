package main

import (
	"errors"
	"fmt"
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

		// only read if it's a file
		if info.Mode().IsDir() {
			fmt.Printf("filesearch: %s is a directory\n", path)
			return nil
		} 
		
		ignore := toIgnore(info.Name())
		if ignore {
			return nil
		}

		if info.Mode().IsRegular() {

			if lin, col, content, err := searchQueryFound(path); err == nil {
				results = append(results, &SearchResult{lin, col, path, content})
			}

			return nil
		} 

		errMsg := fmt.Sprintf("not a regular file: %s", path)
		return errors.New(errMsg)
	}

	if *RecursiveSearch {
		if err := filepath.Walk(path, processFileFn); err != nil {
			return nil, err
		}
	} else {
		info, err := os.Stat(path)
		processFileFn(path, info, err)
	}

	return results, nil
}

