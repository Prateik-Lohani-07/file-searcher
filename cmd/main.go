package main

import (
	"flag"
)


func main() {
	flag.Parse()
	dir, _ := *Dir, *Query
	searchDir(dir)
}