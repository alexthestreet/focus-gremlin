package config

import (
	"path/filepath"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	if got, want := cfg.IntervalMinutes, 60; got != want {
		t.Fatalf("expected interval %d, got %d", want, got)
	}

	if got, want := cfg.ActiveStart, "09:00"; got != want {
		t.Fatalf("expected active start %q, got %q", want, got)
	}

	if got, want := cfg.ActiveEnd, "17:00"; got != want {
		t.Fatalf("expected active end %q, got %q", want, got)
	}

	if got, want := cfg.SnoozeMinutes, 10; got != want {
		t.Fatalf("expected snooze %d, got %d", want, got)
	}

	if cfg.PromptTimeoutSeconds != 0 {
		t.Fatalf("expected prompt timeout disabled, got %d", cfg.PromptTimeoutSeconds)
	}

	if got, want := cfg.Statuses, []string{"on_track", "off_track", "deep_in_the_void", "snooze"}; len(got) != len(want) {
		t.Fatalf("expected %d statuses, got %d", len(want), len(got))
	} else {
		for i := range want {
			if got[i] != want[i] {
				t.Fatalf("expected status %d to be %q, got %q", i, want[i], got[i])
			}
		}
	}
}

func TestSaveLoadRoundTrip(t *testing.T) {
	path := filepath.Join(t.TempDir(), "nested", "config.json")
	want := DefaultConfig()
	want.IntervalMinutes = 45
	want.ActiveStart = "08:30"
	want.TerminalCommand = []string{"kitty", "--hold"}

	if err := Save(path, want); err != nil {
		t.Fatalf("save config: %v", err)
	}

	got, err := Load(path)
	if err != nil {
		t.Fatalf("load config: %v", err)
	}

	if got.IntervalMinutes != want.IntervalMinutes {
		t.Fatalf("expected interval %d, got %d", want.IntervalMinutes, got.IntervalMinutes)
	}

	if got.ActiveStart != want.ActiveStart {
		t.Fatalf("expected active start %q, got %q", want.ActiveStart, got.ActiveStart)
	}

	if len(got.TerminalCommand) != len(want.TerminalCommand) {
		t.Fatalf("expected terminal command length %d, got %d", len(want.TerminalCommand), len(got.TerminalCommand))
	}

	for i := range want.TerminalCommand {
		if got.TerminalCommand[i] != want.TerminalCommand[i] {
			t.Fatalf("expected terminal command %d to be %q, got %q", i, want.TerminalCommand[i], got.TerminalCommand[i])
		}
	}
}
