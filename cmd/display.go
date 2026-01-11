package main

import (
	"fmt"

	"github.com/fatih/color"
)

func displayResults(results []*SearchResult) {
	if len(results) == 0 {
		fmt.Printf("No results found!")
		return
	}

	queryLen := len(Query)

	for _, result := range results {
		displayQueryLocation(result)

		// getting content and surrounding indices of query for highlighting
		content := result.lineContent
		start, end := result.colNum, result.colNum + queryLen

		// getting the various parts of string to highlight the query only
		before, highlightedQuery, after := content[:start], content[start:end], content[end:]

		var display string = before + color.RedString(highlightedQuery) + after
		fmt.Printf("%s: %s\n", result.path, display)
	}
}

func displayQueryLocation(result *SearchResult) {
	showLineNum := *ShowLineNum

	if showLineNum {
		var toPrint string = fmt.Sprintf("[%d,%d]:", result.linNum, result.colNum)
		fmt.Print(color.YellowString(toPrint))
	}
}