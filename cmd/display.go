package main

import (
	"fmt"
	"strconv"
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
	showColNum := *ShowColNum

	var toPrint string = "["

	if showLineNum {
		toPrint += strconv.Itoa(result.linNum)
	} 
	if showColNum {
		toPrint += "," + strconv.Itoa(result.colNum)
	}
	
	toPrint += "]:"

	fmt.Print(toPrint)
}