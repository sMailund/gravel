package main

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
