package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/alexthestreet/focus-gremlin/internal/config"
)

func NewConfigCommand() *Command {
	command := &Command{
		Use:   "config",
		Short: "Manage focus-gremlin configuration",
	}

	command.AddCommand(&Command{
		Use:   "init",
		Short: "Write the default configuration file",
		RunE: func(args []string) error {
			path, err := config.ConfigPath()
			if err != nil {
				return err
			}

			return config.Save(path, config.DefaultConfig())
		},
	}, &Command{
		Use:   "show",
		Short: "Print the resolved configuration",
		RunE: func(args []string) error {
			path, err := config.ConfigPath()
			if err != nil {
				return err
			}

			cfg, err := config.Load(path)
			if err != nil {
				return err
			}

			data, err := json.MarshalIndent(cfg, "", "  ")
			if err != nil {
				return fmt.Errorf("encode config for display: %w", err)
			}

			_, err = os.Stdout.Write(append(data, '\n'))
			return err
		},
	})

	return command
}
