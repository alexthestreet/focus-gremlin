package integration

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/alexthestreet/focus-gremlin/internal/cli"
	"github.com/alexthestreet/focus-gremlin/internal/config"
	runtimestate "github.com/alexthestreet/focus-gremlin/internal/runtime"
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

func TestPromptSnoozeUpdatesRuntimeState(t *testing.T) {
	root := t.TempDir()
	configHome := filepath.Join(root, "config")
	dataPath := filepath.Join(root, "history.db")
	runtimeDir := filepath.Join(root, "runtime")

	t.Setenv("XDG_CONFIG_HOME", configHome)
	t.Setenv("FOCUS_GREMLIN_DATA_PATH", dataPath)
	t.Setenv("FOCUS_GREMLIN_PROMPT_CHOICE", "snooze")
	t.Setenv("XDG_RUNTIME_DIR", runtimeDir)

	configPath, err := config.ConfigPath()
	if err != nil {
		t.Fatalf("config path: %v", err)
	}
	cfg := config.DefaultConfig()
	cfg.SnoozeMinutes = 10
	if err := config.Save(configPath, cfg); err != nil {
		t.Fatalf("save config: %v", err)
	}

	command := cli.NewPromptCommand()
	before := time.Now().UTC()
	if err := command.RunE(nil); err != nil {
		t.Fatalf("run prompt command: %v", err)
	}
	after := time.Now().UTC()

	statePath := filepath.Join(runtimeDir, "focus-gremlin", "state.json")
	state, err := runtimestate.LoadState(statePath)
	if err != nil {
		t.Fatalf("load state: %v", err)
	}

	min := before.Add(10 * time.Minute)
	max := after.Add(10 * time.Minute)
	if state.SnoozedUntil.Before(min) || state.SnoozedUntil.After(max) {
		t.Fatalf("expected snoozed_until between %s and %s, got %s", min, max, state.SnoozedUntil)
	}

	if state.PromptActive {
		t.Fatal("expected prompt_active to be false after prompt completion")
	}
}
