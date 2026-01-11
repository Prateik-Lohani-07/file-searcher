package cli

import (
	"errors"
	"flag"
	"fmt"
	"path/filepath"
)

func ParseArgs() (string, []string, error)  {
	flag.Parse()

	if (*Help) {
		displayHelp()
		return "", nil, nil
	}

	args := flag.Args()

	if len(args) < 2 {
		return "", nil, errors.New("must pass in at least 2 args (query and file pattern)")
	}

	query, glob := args[0], args[1:]
	paths := []string{}

	for _, g := range glob {
		expandedGlob, err := expandGlob(g)

		if err != nil {
			errMsg := fmt.Sprintf("error while parsing glob pattern: %v", err)
			return "", nil, errors.New(errMsg)
		}

		paths = append(paths, expandedGlob...)
	}

	return query, paths, nil
}

func expandGlob(glob string) ([]string, error) {
	matches, err := filepath.Glob(glob)
	if err != nil {
		return nil, err
	}

	return matches, nil
}