package cli

import (
	"fmt"

	"github.com/alexthestreet/focus-gremlin/internal/config"
	"github.com/alexthestreet/focus-gremlin/internal/launcher"
)

func NewDaemonCommand() *Command {
	return &Command{
		Use:   "daemon",
		Short: "Run the focus scheduler daemon",
		RunE: func(args []string) error {
			cfg, err := loadDaemonConfig()
			if err != nil {
				return err
			}

			_, err = launcher.BuildCommand(cfg.TerminalCommand, []string{"focus-gremlin", "prompt"})
			if err != nil {
				return fmt.Errorf("prepare prompt launcher: %w", err)
			}

			return nil
		},
	}
}

func loadDaemonConfig() (config.Config, error) {
	path, err := config.ConfigPath()
	if err != nil {
		return config.Config{}, err
	}

	return config.Load(path)
}
