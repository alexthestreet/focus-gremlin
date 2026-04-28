package cli

func NewDaemonCommand() *Command {
	return &Command{
		Use:   "daemon",
		Short: "Run the focus scheduler daemon",
		RunE: func(args []string) error {
			return nil
		},
	}
}
