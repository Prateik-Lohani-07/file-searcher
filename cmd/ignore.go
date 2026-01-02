package main

import "path/filepath"

var ignore = map[string]struct{}{
	".git": {},
	".env": {},
}

func toIgnore(path string) bool {
	base := filepath.Base(path)
	_, exists := ignore[base]
	return exists
}