package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		usage()
		return
	}

	if strings.Compare(args[0], "-s") == 0 {
		search(args[1])
	} else {
		usage()
	}
}

func usage() {
	println("unrecognized command")
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

	if len(matches) > 0 {
		for _, match := range matches {
			fmt.Printf("%v -> %v\n", match.lineNumber, match.text)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

type match struct {
	text       string
	lineNumber int
}
