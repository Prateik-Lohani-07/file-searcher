package main

import (
	"flag"
)

var (
	RecursiveSearch = flag.Bool("r", false, "Specifies whether recursive directory search is to be performed.")
	ShowLineNum = flag.Bool("n", false, "Whether to show the line and column numbers where the query occurs in the file.")
	Help = flag.Bool("help", false, "Show the various flags, their meaning, and their usage")
)