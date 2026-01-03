package main

import (
	// "fmt"
	"flag"
)


func main() {
	flag.Parse()
	searchResults := searchDir()
	displayResults(searchResults)
}