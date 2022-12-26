package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

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
