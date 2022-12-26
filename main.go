package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var commands = []Command{
	&UsageCommand{},
	&LinksCommand{},
	&SearchCommand{},
}

func main() {
	args := os.Args[1:]

	cmd := getCommand(args)
	cmd.execute()
}

func getCommand(args []string) Command {
	if len(args) == 0 {
		usageCommand := UsageCommand{}
		return usageCommand.create(args)
	}

	flag := args[0]
	for _, cmd := range commands {
		if strings.Compare(flag, cmd.flag()) == 0 {
			return cmd.create(args)
		}
	}

	return &UsageCommand{}
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

func findLinks(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic("")
	}
	defer f.Close()

	linkRegex := regexp.MustCompile("\\[\\S+\\]")

	var links []string
	scanner := bufio.NewScanner(f)
	line := 1
	for scanner.Scan() {
		text := scanner.Text()
		matches := linkRegex.FindAllString(text, -1)
		links = append(links, matches...)

		line++
	}

	links = removeDuplicateStr(links)

	for _, link := range links {
		fmt.Println(link)
	}
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

type match struct {
	text       string
	lineNumber int
}

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

type UsageCommand struct {
}

func (c *UsageCommand) create(args []string) Command {
	return &UsageCommand{}
}

func (c *UsageCommand) flag() string {
	return "?"
}

func (c *UsageCommand) description() string {
	return "show list of commands"
}

func (c *UsageCommand) execute() {
	fmt.Println("gravel - local-only obsidian alternative for the Command line")

	for _, command := range commands {
		fmt.Printf("\t%v %v\n", command.flag(), command.description())
	}
}

type LinksCommand struct {
	path string
}

func (c *LinksCommand) create(args []string) Command {
	return &LinksCommand{path: args[1]}
}

func (c *LinksCommand) flag() string {
	return "-l"
}

func (c *LinksCommand) description() string {
	return "show all links in given file"
}

func (c *LinksCommand) execute() {
	findLinks(c.path)
}
