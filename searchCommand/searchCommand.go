package searchCommand

import (
	"grvl/command"
	"grvl/search"
)

type SearchCommand struct {
	searchTerm string
}

func (c *SearchCommand) Create(args []string) command.Command {
	return &SearchCommand{searchTerm: args[1]}
}

func (c *SearchCommand) Flag() string {
	return "-s"
}

func (c *SearchCommand) Description() string {
	return "search all files for a file containing search term"
}

func (c *SearchCommand) Execute() {
	search.SearchAllFiles(c.searchTerm)
}
