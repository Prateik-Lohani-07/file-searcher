package main

import "path/filepath"

var ignore = map[string]struct{}{
	".git": {},
	".exe": {},
	".jpeg": {},
	".jpg": {},
	".png": {},
	".mp4": {},
	".mkv": {},
}

func toIgnore(path string) bool {
	var toIgnore bool = false

	// ignore if base path to be ignored
	base := filepath.Base(path)
	_, exists := ignore[base]
	
	// ignore if kind of file is to be ignored
	fileExtension := filepath.Ext(path)
	_, extToIgnore := ignore[fileExtension]

	toIgnore = toIgnore || exists || extToIgnore
	return toIgnore
}