package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexthestreet/focus-gremlin/internal/config"
)

func NewStatusCommand() *Command {
	return &Command{
		Use:   "status",
		Short: "Show runtime and storage status",
		RunE: func(args []string) error {
			configPath, err := config.ConfigPath()
			if err != nil {
				return err
			}

			dataPath, err := config.DataPath()
			if err != nil {
				return err
			}

			runtimeDir, err := config.RuntimeDir()
			if err != nil {
				return err
			}

			statePath := filepath.Join(runtimeDir, "state.json")
			lockPath := filepath.Join(runtimeDir, "daemon.lock")

			fmt.Fprintf(os.Stdout, "config: %s\n", configPath)
			fmt.Fprintf(os.Stdout, "data: %s\n", dataPath)
			fmt.Fprintf(os.Stdout, "runtime: %s\n", runtimeDir)
			fmt.Fprintf(os.Stdout, "state file present: %t\n", fileExists(statePath))
			fmt.Fprintf(os.Stdout, "lock file present: %t\n", fileExists(lockPath))
			return nil
		},
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
