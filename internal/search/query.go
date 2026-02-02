package search

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strings"
	"sync"

	"filesearch/internal/cli"
	"filesearch/internal/models"

	"github.com/sahilm/fuzzy"
)

type SearchResult = models.SearchResult

func findQueryInFile(path string, limits chan bool, results chan<- SearchResult, wg *sync.WaitGroup) {
	defer wg.Done()

	limits <- true

	defer func(){ <-limits }()

	search(path, results)
}

func search(path string, results chan<- SearchResult) {
	f := *cli.Fuzzy

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open file %s: %v", path, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if f {
		fuzzySearch(path, lines, results)
	} else {
		strictMatch(path, lines, results)
	}
}

func fuzzySearch(path string, lines []string, results chan<- SearchResult) {
	query := Query
	matches := fuzzy.Find(query, lines)

	maxAbs := math.MinInt
	minVal := math.MaxInt

	for _, m := range matches {
		absScore := math.Abs(float64(m.Score))
		maxAbs = max(maxAbs, int(absScore))
		minVal = min(minVal, m.Score)
	}

	// normalizing the results scores and adding ones with score greater than 0 -> TODO: experiment with this
	for idx, m := range matches {
		normalizedScore := (float64(m.Score) - float64(minVal)) / float64(maxAbs)

		if normalizedScore <= 0.5 {
			break
		}

		results <- SearchResult{
			LinNum:      idx,
			ColNum:      m.MatchedIndexes[0], // NOTE: matched indexes gives 1-indexed col number
			Path:        path,
			LineContent: m.Str,
		}
	}
}

func strictMatch(path string, lines []string, results chan<- SearchResult) {
	query := Query
	w := *cli.Word

	for idx, line := range lines {
		if !w {
			if colNum := strings.Index(line, Query); colNum != -1 {
				results <- SearchResult{
					LinNum:      idx,
					ColNum:      colNum, // NOTE: matched indexes gives 1-indexed col number
					Path:        path,
					LineContent: line,
				}
			}

			// perform strict word search
		} else {
			words := strings.Split(line, " ")

			if idx := slices.Index(words, query); idx != -1 {
				results <- SearchResult{
					LinNum:      idx,
					ColNum:      getWordColNum(idx, words), // NOTE: matched indexes gives 1-indexed col number
					Path:        path,
					LineContent: line,
				}
			}
		}
	}
}

// this function assumes that the word has indeed been found and that idxInWordSlice != -1
func getWordColNum(idxInWordSlice int, words []string) int {
	before := words[:idxInWordSlice]

	colNum := 0
	for _, w := range before {
		colNum += len(w) + 1 // 1 for the whitespace
	}

	return colNum
}
