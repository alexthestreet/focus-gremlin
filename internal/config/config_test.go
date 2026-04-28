package config

import "testing"

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
