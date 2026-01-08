package main

import (
	"path/filepath"
)

var ignore = map[string]bool{
	".git": true,
	".exe": true,
	".jpeg": true,
	".jpg": true,
	".png": true,
	".mp4": true,
	".mkv": true,
}

func toIgnore(path string) bool {
	var toIgnore bool = false

	// ignore if base path to be ignored
	base := filepath.Base(path)
	exists := ignore[base]
	
	// ignore if kind of file is to be ignored
	fileExtension := filepath.Ext(path)
	_, extToIgnore := ignore[fileExtension]

	toIgnore = toIgnore || exists || extToIgnore
	return toIgnore
}