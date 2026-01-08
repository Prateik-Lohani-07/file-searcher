package main

import (
	"path/filepath"
)

func expandGlob(glob string) ([]string, error) {
	matches, err := filepath.Glob(glob)
	if err != nil {
		return nil, err
	}

	return matches, nil
}