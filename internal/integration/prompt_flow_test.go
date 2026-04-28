package integration

import (
	"path/filepath"
	"testing"

	"github.com/alexthestreet/focus-gremlin/internal/cli"
	"github.com/alexthestreet/focus-gremlin/internal/config"
	"github.com/alexthestreet/focus-gremlin/internal/storage"
)

func TestPromptSubmissionCreatesEvent(t *testing.T) {
	root := t.TempDir()
	configHome := filepath.Join(root, "config")
	dataPath := filepath.Join(root, "history.db")

	t.Setenv("XDG_CONFIG_HOME", configHome)
	t.Setenv("FOCUS_GREMLIN_DATA_PATH", dataPath)
	t.Setenv("FOCUS_GREMLIN_PROMPT_CHOICE", "on_track")

	configPath, err := config.ConfigPath()
	if err != nil {
		t.Fatalf("config path: %v", err)
	}
	if err := config.Save(configPath, config.DefaultConfig()); err != nil {
		t.Fatalf("save config: %v", err)
	}

	command := cli.NewPromptCommand()
	if err := command.RunE(nil); err != nil {
		t.Fatalf("run prompt command: %v", err)
	}

	store, err := storage.Open(dataPath)
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	defer store.Close()

	events, err := store.RecentEvents(1)
	if err != nil {
		t.Fatalf("read events: %v", err)
	}

	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}

	if got, want := events[0].Status, "on_track"; got != want {
		t.Fatalf("expected status %q, got %q", want, got)
	}
}
