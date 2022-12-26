package main

import (
	"bufio"
	"fmt"
	"os"
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
