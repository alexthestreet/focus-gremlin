package storage

import (
	"path/filepath"
	"testing"
	"time"
)

func TestInsertEvent(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "history.db")
	store, err := Open(dbPath)
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	defer store.Close()

	want := Event{
		Status:     "on_track",
		RecordedAt: time.Date(2026, time.April, 28, 9, 0, 0, 0, time.UTC),
	}
	if err := store.InsertEvent(want); err != nil {
		t.Fatalf("insert event: %v", err)
	}

	events, err := store.RecentEvents(1)
	if err != nil {
		t.Fatalf("read recent events: %v", err)
	}

	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}

	got := events[0]
	if got.Status != want.Status {
		t.Fatalf("expected status %q, got %q", want.Status, got.Status)
	}

	if !got.RecordedAt.Equal(want.RecordedAt) {
		t.Fatalf("expected recorded_at %s, got %s", want.RecordedAt, got.RecordedAt)
	}
}
