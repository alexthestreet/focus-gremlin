package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/alexthestreet/focus-gremlin/internal/checkin"
	"github.com/alexthestreet/focus-gremlin/internal/config"
	runtimestate "github.com/alexthestreet/focus-gremlin/internal/runtime"
	"github.com/alexthestreet/focus-gremlin/internal/storage"
	prompttui "github.com/alexthestreet/focus-gremlin/internal/tui/prompt"
	tea "github.com/charmbracelet/bubbletea"
)

func NewPromptCommand() *Command {
	return &Command{
		Use:   "prompt",
		Short: "Run the interactive focus prompt",
		RunE: func(args []string) error {
			cfg, err := loadPromptConfig()
			if err != nil {
				return err
			}

			choice := os.Getenv("FOCUS_GREMLIN_PROMPT_CHOICE")
			if choice == "" {
				model := prompttui.NewModel(cfg.Statuses)
				program := tea.NewProgram(model)
				finalModel, err := program.Run()
				if err != nil {
					return fmt.Errorf("run prompt: %w", err)
				}

				resolved, ok := finalModel.(prompttui.Model)
				if !ok {
					return fmt.Errorf("unexpected prompt model type %T", finalModel)
				}

				choice = resolved.SelectedChoice()
			}

			store, err := storage.Open(promptDataPath())
			if err != nil {
				return err
			}
			defer store.Close()

			now := time.Now().UTC()
			if err := checkin.SubmitResponse(store, checkin.Result{
				Status:     choice,
				RecordedAt: now,
			}); err != nil {
				return err
			}

			return updatePromptState(choice, cfg, now)
		},
	}
}

func loadPromptConfig() (config.Config, error) {
	path, err := config.ConfigPath()
	if err != nil {
		return config.Config{}, err
	}

	return config.Load(path)
}

func promptDataPath() string {
	if value := os.Getenv("FOCUS_GREMLIN_DATA_PATH"); value != "" {
		return value
	}

	path, err := config.DataPath()
	if err != nil {
		return "history.db"
	}

	return path
}

func updatePromptState(choice string, cfg config.Config, now time.Time) error {
	statePath, err := promptStatePath()
	if err != nil {
		return nil
	}

	state, err := runtimestate.LoadState(statePath)
	if err != nil {
		return err
	}

	state.PromptActive = false
	state.SnoozedUntil = time.Time{}
	if choice == "snooze" {
		state.SnoozedUntil = now.Add(time.Duration(cfg.SnoozeMinutes) * time.Minute)
	}

	return runtimestate.SaveState(statePath, state)
}

func promptStatePath() (string, error) {
	if value := os.Getenv("FOCUS_GREMLIN_STATE_PATH"); value != "" {
		return value, nil
	}

	runtimeDir, err := config.RuntimeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(runtimeDir, "state.json"), nil
}
