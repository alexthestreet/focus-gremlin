package scheduler

import (
	"testing"
	"time"
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
