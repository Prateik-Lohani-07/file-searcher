package main

import (
	"fmt"
)

func displayResults(results []*SearchResult) {
	if len(results) == 0 {
		fmt.Printf("No results found!")
		return
	}

	for _, result := range results {
		displayQueryLocation(result)

		fmt.Printf("%s:\t%s\n", result.path, result.lineContent)
	}
}

func displayQueryLocation(result *SearchResult) {
	showLineNum := *ShowLineNum

	if showLineNum {
		var toPrint string = fmt.Sprintf("[%d,%d]:", result.linNum, result.colNum)
		fmt.Print(toPrint)
	}
}