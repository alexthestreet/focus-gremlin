package cli

import (
	"errors"
	"fmt"
	"strings"
)

type Command struct {
	Use      string
	Short    string
	RunE     func(args []string) error
	commands []*Command
}

func NewRootCommand() *Command {
	root := &Command{
		Use:   "focus-gremlin",
		Short: "A periodic focus check-in companion",
	}

	root.AddCommand(NewDaemonCommand(), NewPromptCommand(), NewConfigCommand())

	return root
}

func (c *Command) AddCommand(commands ...*Command) {
	c.commands = append(c.commands, commands...)
}

func (c *Command) Commands() []*Command {
	return append([]*Command(nil), c.commands...)
}

func (c *Command) Execute(args []string) error {
	if len(args) == 0 {
		if c.RunE != nil {
			return c.RunE(nil)
		}

		return nil
	}

	next, remainder, ok := c.matchSubcommand(args)
	if !ok {
		return fmt.Errorf("unknown command: %s", strings.Join(args, " "))
	}

	if next == nil {
		return errors.New("command is nil")
	}

	return next.Execute(remainder)
}

func (c *Command) matchSubcommand(args []string) (*Command, []string, bool) {
	name := args[0]
	for _, child := range c.commands {
		if child.Use == name {
			return child, args[1:], true
		}
	}

	return nil, nil, false
}
