package search

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func SearchAllFiles(term string, printer ResultPrinter) {
	e := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		matches := searchFile(path, term)
		printer.print(info, matches)
		return nil
	})
	if e != nil {
		log.Fatal(e)
	}
}

func searchFile(path string, term string) []match {
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
		return nil
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

	return matches
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

type match struct {
	text       string
	lineNumber int
}
