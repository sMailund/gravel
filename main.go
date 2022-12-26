package main

import (
	"fmt"
	"os"
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
