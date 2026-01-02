package main

import (
	"flag"
)

var (
	Dir = flag.Int("dir", 1, "specifies the directory in which search operation is to be performed")
	Query = flag.Int("query", 2, "specifies the keyword that is to be searched for in the search directory")
) 