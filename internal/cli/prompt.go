package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/alexthestreet/focus-gremlin/internal/checkin"
	"github.com/alexthestreet/focus-gremlin/internal/config"
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

			return checkin.SubmitResponse(store, checkin.Result{
				Status:     choice,
				RecordedAt: time.Now().UTC(),
			})
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
