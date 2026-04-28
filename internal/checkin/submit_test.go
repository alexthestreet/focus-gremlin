package checkin

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/alexthestreet/focus-gremlin/internal/storage"
)

func TestSubmitResponse(t *testing.T) {
	store, err := storage.Open(filepath.Join(t.TempDir(), "history.db"))
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	defer store.Close()

	result := Result{
		Status:     "off_track",
		RecordedAt: time.Date(2026, time.April, 28, 12, 0, 0, 0, time.UTC),
	}
	if err := SubmitResponse(store, result); err != nil {
		t.Fatalf("submit response: %v", err)
	}

	events, err := store.RecentEvents(1)
	if err != nil {
		t.Fatalf("read events: %v", err)
	}

	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}

	if got, want := events[0].Status, result.Status; got != want {
		t.Fatalf("expected status %q, got %q", want, got)
	}
}
