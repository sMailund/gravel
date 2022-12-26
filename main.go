package main

import (
	"fmt"
	"grvl/command"
	"grvl/linksCommand"
	"grvl/searchCommand"
	"os"
	"strings"
)

var commands = []command.Command{
	&UsageCommand{},
	&linksCommand.LinksCommand{},
	&searchCommand.SearchCommand{},
}

func main() {
	args := os.Args[1:]

	cmd := getCommand(args)
	cmd.Execute()
}

func getCommand(args []string) command.Command {
	if len(args) == 0 {
		usageCommand := UsageCommand{}
		return usageCommand.Create(args)
	}

	argument := args[0]
	for _, cmd := range commands {
		if strings.Compare(argument, cmd.Flag()) == 0 || strings.Compare(argument, cmd.Keyword()) == 0 {
			return cmd.Create(args)
		}
	}

	return &UsageCommand{}
}

type UsageCommand struct {
}

func (c *UsageCommand) Keyword() string {
	return "usage"
}

func (c *UsageCommand) Create(args []string) command.Command {
	return &UsageCommand{}
}

func (c *UsageCommand) Flag() string {
	return "?"
}

func (c *UsageCommand) Description() string {
	return "show list of commands"
}

func (c *UsageCommand) Execute() {
	fmt.Println("gravel - local-only obsidian alternative for the Command line")

	for _, command := range commands {
		fmt.Printf("\t%v, %v \t-> %v\n", command.Keyword(), command.Flag(), command.Description())
	}
}
