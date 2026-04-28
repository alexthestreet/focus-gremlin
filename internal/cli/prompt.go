package cli

func NewPromptCommand() *Command {
	return &Command{
		Use:   "prompt",
		Short: "Run the interactive focus prompt",
		RunE: func(args []string) error {
			return nil
		},
	}
}
