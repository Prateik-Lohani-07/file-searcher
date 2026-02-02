package search

import (
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"filesearch/internal/cli"
)

var Query string

func SearchAllPaths(query string, paths []string) ([]SearchResult, error) {
	results := make(chan SearchResult, 100) // arbitrary buffer size
	Query = query

	resultsLst := []SearchResult{}
	go func() {
		for r := range results {
			resultsLst = append(resultsLst, r)
		}
	}()

	for _, path := range paths {
		err := startSearchWorkers(path, results)
		if err != nil {
			return nil, err
		}
	}

	close(results)
	return resultsLst, nil
}

func startSearchWorkers(dir string, results chan<- SearchResult) error {
	workers := 2 * runtime.GOMAXPROCS(0)

	limits := make(chan bool, workers)
	wg := new(sync.WaitGroup)

	wg.Add(1)
	err := searchTree(dir, limits, results, wg)

	if err != nil {
		return err
	}

	wg.Wait()
	return nil
}

func searchTree(dir string, limits chan bool, results chan<- SearchResult, wg *sync.WaitGroup) error {
	defer wg.Done()

	processFileFn := func(path string, info fs.FileInfo, err error) error {
		if err != nil && err != os.ErrNotExist {
			return err;
		}

		ignore := toIgnore(info.Name())
		if ignore {
			return filepath.SkipDir
		}

		if info.Mode().IsDir() && path != dir {
			wg.Add(1)
			go searchTree(path, limits, results, wg)
			
			return filepath.SkipDir

		} else {
			wg.Add(1)
			go findQueryInFile(path, limits, results, wg)
		}

		return nil
	}

	limits <- true
	defer func(){ <-limits }()

	if *cli.RecursiveSearch {
		if err := filepath.Walk(dir, processFileFn); err != nil {
			return err
		}
	} else {
		info, err := os.Stat(dir)
		err = processFileFn(dir, info, err)

		if err == filepath.SkipDir || err == filepath.SkipAll {
			return nil
		}
		if err != nil {
			return err
		}
	}

	return nil
}
