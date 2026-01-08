package main

import (
	"os"
	"bufio"
	"strings"
	"log"
	"errors"
)

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

		if charNum := strings.Index(line, Query); charNum != -1 {
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