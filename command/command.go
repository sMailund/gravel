package command

type Command interface {
	Create(args []string) Command
	Execute()
	Flag() string
	Keyword() string
	Description() string
}
