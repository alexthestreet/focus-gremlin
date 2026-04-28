package scheduler

import (
	"testing"
	"time"

	runtimestate "github.com/alexthestreet/focus-gremlin/internal/runtime"
)

func TestNextCheckInWithinActiveHours(t *testing.T) {
	now := time.Date(2026, time.April, 28, 10, 0, 0, 0, time.UTC)
	next, err := NextCheckIn(now, Options{
		Interval:    time.Hour,
		ActiveStart: "09:00",
		ActiveEnd:   "17:00",
	})
	if err != nil {
		t.Fatalf("next check-in: %v", err)
	}

	want := time.Date(2026, time.April, 28, 11, 0, 0, 0, time.UTC)
	if !next.Equal(want) {
		t.Fatalf("expected %s, got %s", want, next)
	}
}

func TestNextCheckInAfterActiveHours(t *testing.T) {
	now := time.Date(2026, time.April, 28, 16, 30, 0, 0, time.UTC)
	next, err := NextCheckIn(now, Options{
		Interval:    time.Hour,
		ActiveStart: "09:00",
		ActiveEnd:   "17:00",
	})
	if err != nil {
		t.Fatalf("next check-in: %v", err)
	}

	want := time.Date(2026, time.April, 29, 9, 0, 0, 0, time.UTC)
	if !next.Equal(want) {
		t.Fatalf("expected %s, got %s", want, next)
	}
}

func TestSnoozeDelaysNextPrompt(t *testing.T) {
	now := time.Date(2026, time.April, 28, 10, 0, 0, 0, time.UTC)
	next, err := NextAction(now, Options{
		Interval:    time.Hour,
		ActiveStart: "09:00",
		ActiveEnd:   "17:00",
	}, runtimestate.State{
		SnoozedUntil: time.Date(2026, time.April, 28, 10, 10, 0, 0, time.UTC),
	})
	if err != nil {
		t.Fatalf("next action: %v", err)
	}

	if next.SkipPrompt {
		t.Fatal("expected prompt to remain eligible after snooze")
	}

	want := time.Date(2026, time.April, 28, 10, 10, 0, 0, time.UTC)
	if !next.When.Equal(want) {
		t.Fatalf("expected %s, got %s", want, next.When)
	}
}

func TestActivePromptBlocksLaunch(t *testing.T) {
	now := time.Date(2026, time.April, 28, 10, 0, 0, 0, time.UTC)
	next, err := NextAction(now, Options{
		Interval:    time.Hour,
		ActiveStart: "09:00",
		ActiveEnd:   "17:00",
	}, runtimestate.State{PromptActive: true})
	if err != nil {
		t.Fatalf("next action: %v", err)
	}

	if !next.SkipPrompt {
		t.Fatal("expected active prompt to block a new prompt launch")
	}

	if !next.When.Equal(now) {
		t.Fatalf("expected blocked prompt to keep current time %s, got %s", now, next.When)
	}
}
