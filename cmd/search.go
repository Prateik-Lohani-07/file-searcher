package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

func getFiles(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if toIgnore(info.Name()) {
		if info.IsDir() {
			return filepath.SkipDir
		}

		return nil
	}

	fmt.Println(info.Name())
	return nil
}

func searchDir(dir string) {
	err := filepath.Walk(dir, getFiles)
	if err != nil {
		log.Fatal(err)
	}
}