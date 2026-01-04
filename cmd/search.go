package main

import (
	"bufio"
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)


func searchDir() ([]*SearchResult, error) {
	dir := *Dir
	var results = []*SearchResult{}

	walkFn := func (path string, info fs.FileInfo, err error) error {
		if err != nil {
			if info != nil && info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		ignore := toIgnore(info.Name())
		
		if ignore {
			if info.IsDir() {
				return filepath.SkipDir
			}

			return nil
		}

		// adding the line column 
		if !info.IsDir() {
			if lin, col, content, err := searchQueryFound(path); err == nil {
				results = append(results, &SearchResult{lin, col, path, content})
			}
		}

		return nil
	}
	
	if err := filepath.Walk(dir, walkFn); err != nil {
		return nil, err
	}

	return results, nil
}

func searchQueryFound(path string) (int, int, string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open file %s: %v", path, err)
		return -1, -1, "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	linNum, colNum := -1, -1
	lineContent := ""

	for scanner.Scan() {
		linNum++
		line := scanner.Text()

		if charNum := strings.Index(line, *Query); charNum != -1 {
			lineContent = line
			colNum = charNum
			break
		}
	}

	if err := scanner.Err(); err != nil {
        return -1, -1, "", err
    }
	
	if colNum == -1 {
		return -1, -1, "", errors.New("query not found!")
	}

	return linNum + 1, colNum + 1, lineContent, nil
}