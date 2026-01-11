package cli

import (
	"fmt"

	"filesearch/internal/models"

	"github.com/fatih/color"
)

type SearchResult = models.SearchResult

func DisplayResults(query string, results []SearchResult) {
	if len(results) == 0 {
		fmt.Printf("No results found!")
		return
	}

	queryLen := len(query)

	for _, result := range results {
		displayQueryLocation(result)

		// getting content and surrounding indices of query for highlighting
		content := result.LineContent
		start, end := result.ColNum, result.ColNum + queryLen

		// getting the various parts of string to highlight the query only
		before, highlightedQuery, after := content[:start], content[start:end], content[end:]

		var display string = before + color.RedString(highlightedQuery) + after
		fmt.Printf("%s: %s\n", result.Path, display)
	}
}

func displayQueryLocation(result SearchResult) {
	showLineNum := *ShowLineNum

	if showLineNum {
		var toPrint string = fmt.Sprintf("[%d,%d]:", result.LinNum, result.ColNum)
		fmt.Print(color.YellowString(toPrint))
	}
}