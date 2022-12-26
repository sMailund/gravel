package main

type Command interface {
	create(args []string) Command
	execute()
	flag() string
	description() string
}
