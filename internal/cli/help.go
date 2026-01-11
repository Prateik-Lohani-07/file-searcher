package cli

import "fmt"

func displayHelp() {
	msg := `
FileSearcher - A simple CLI tool to search for keywords in files

Usage:
  filesearch [options] <query> <file patterns...>

Positional arguments:
  query            The keyword or string to search for.
  file patterns    One or more file patterns or paths (e.g., *.go, ./src).

Options:
  -r                 Enable recursive search through subdirectories (default: false)
  -n                 Show the line and column numbers where the query occurs in the file
  -fz				 Perform fuzzy search
  -w				 Search for words (whitespace separated)
  -help              Show this help message

Examples:
  filesearch -w noOfRequests src/**/*.js
  filesearch -r func *.go
  filesearch -n TODO .
  filesearch -r -n error *.txt
  filesearch -fz cat drive/*

Notes:
- Ignored directories and files (e.g., .git, .env) are skipped automatically.
- File patterns support *, ? (globs); use -r for recursive search.`

	fmt.Println(msg)
}