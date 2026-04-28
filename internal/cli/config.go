package cli

func NewConfigCommand() *Command {
	return &Command{
		Use:   "config",
		Short: "Manage focus-gremlin configuration",
		RunE: func(args []string) error {
			return nil
		},
	}
}
