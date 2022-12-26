package linksCommand

import (
	"bufio"
	"fmt"
	"grvl/command"
	"os"
	"regexp"
)

type LinksCommand struct {
	path string
}

func (c *LinksCommand) Create(args []string) command.Command {
	return &LinksCommand{path: args[1]}
}

func (c *LinksCommand) Flag() string {
	return "-l"
}

func (c *LinksCommand) Description() string {
	return "show all links in given file"
}

func (c *LinksCommand) Execute() {
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

	for i, link := range links {
		fmt.Printf("%v -> %v\n", i, link)
	}

	var selection int
	_, err = fmt.Scanf("%d", &selection)

	if err != nil {
		fmt.Println("could not read input.")
		return
	}

	if selection < 0 || selection > len(links)-1 {
		fmt.Println("out of range")
		return
	}

	selected := links[selection]

	fmt.Println(selected)
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
