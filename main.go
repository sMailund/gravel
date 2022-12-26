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

	cmd := getCommand(args)
	cmd.execute()
}

func getCommand(args []string) command {
	if len(args) == 0 {
		return &UsageCommand{}
	} else if strings.Compare(args[0], "-s") == 0 {
		return &SearchCommand{searchTerm: args[1]}
	} else if strings.Compare(args[0], "-l") == 0 {
		return &LinksCommand{path: args[1]}
	} else {
		return &UsageCommand{}
	}
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

type match struct {
	text       string
	lineNumber int
}

type command interface {
	execute()
}

type SearchCommand struct {
	searchTerm string
}

func (c *SearchCommand) execute() {
	search(c.searchTerm)
}

type UsageCommand struct {
}

func (c *UsageCommand) execute() {
	println("unrecognized command")
}

type LinksCommand struct {
	path string
}

func (c *LinksCommand) execute() {
	println(c.path)
}
