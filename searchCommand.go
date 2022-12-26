package main

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
