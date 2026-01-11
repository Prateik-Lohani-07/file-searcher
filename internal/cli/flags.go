package cli

import (
	"flag"
)

var (
	RecursiveSearch = flag.Bool("r", false, "Specifies whether recursive directory search is to be performed.")
	ShowLineNum = flag.Bool("n", false, "Whether to show the line and column numbers where the query occurs in the file.")
	Fuzzy = flag.Bool("fz", false, "Whether to do fuzzy search (NOTE: gives word search only).")
	Word = flag.Bool("w", false, "Whether to search for words rather than regular string matching")
	Help = flag.Bool("help", false, "Show the various flags, their meaning, and their usage.")
)