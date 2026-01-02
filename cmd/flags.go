package main

import (
	"flag"
)

var (
	Dir = flag.String("dir", "", "specifies the directory in which search operation is to be performed")
	Query = flag.String("query", "", "specifies the keyword that is to be searched for in the search directory")
)