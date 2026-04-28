package cli

import (
	"context"
	"path/filepath"
	"time"

	"github.com/alexthestreet/focus-gremlin/internal/config"
	"github.com/alexthestreet/focus-gremlin/internal/daemon"
	"github.com/alexthestreet/focus-gremlin/internal/scheduler"
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

			runtimeDir, err := config.RuntimeDir()
			if err != nil {
				return err
			}

			return daemon.Run(context.Background(), daemon.Options{
				LockPath:  filepath.Join(runtimeDir, "daemon.lock"),
				StatePath: filepath.Join(runtimeDir, "state.json"),
				Scheduler: scheduler.Options{
					Interval:    time.Duration(cfg.IntervalMinutes) * time.Minute,
					ActiveStart: cfg.ActiveStart,
					ActiveEnd:   cfg.ActiveEnd,
				},
				TerminalCommand: cfg.TerminalCommand,
				AppCommand:      []string{"focus-gremlin", "prompt"},
			})
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
