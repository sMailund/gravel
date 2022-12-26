package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type SearchCommand struct {
	searchTerm string
}

func (c *SearchCommand) create(args []string) Command {
	return &SearchCommand{searchTerm: args[1]}
}

func (c *SearchCommand) flag() string {
	return "-s"
}

func (c *SearchCommand) description() string {
	return "search for a file containing search term"
}

func (c *SearchCommand) execute() {
	search(c.searchTerm)
}

func search(term string) {
	e := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		searchFile(path, term)
		return nil
	})
	if e != nil {
		log.Fatal(e)
	}
}

func searchFile(path string, term string) {
	f, err := os.Open(path)
	if err != nil {
		panic("")
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		panic("")
	}

	if fileInfo.IsDir() {
		return
	}

	scanner := bufio.NewScanner(f)

	matches := []match{}

	line := 1
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, term) {
			match := match{
				text:       text,
				lineNumber: line,
			}
			matches = append(matches, match)
		}

		line++
	}

	printMatches(fileInfo, matches)
}

func printMatches(fileInfo os.FileInfo, matches []match) {
	if len(matches) > 0 {
		fmt.Printf("%v\n", fileInfo.Name())
		for _, match := range matches {
			fmt.Printf("%v -> %v\n", match.lineNumber, match.text)
		}
		fmt.Println()
	}
}
