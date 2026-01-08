package main

import (
	"flag"
)

var (
	RecursiveSearch = flag.Bool("r", false, "Whether to search for all subdirectories recursively under the current directory")
	ShowLineNum = flag.Bool("n", false, "Whether to show the line and column numbers where the query occurs in the file.")
)