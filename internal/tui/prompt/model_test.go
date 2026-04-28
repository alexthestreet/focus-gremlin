package prompt

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestDefaultSelection(t *testing.T) {
	model := NewModel([]string{"on_track", "off_track", "deep_in_the_void", "snooze"})
	if got, want := model.SelectedIndex(), 0; got != want {
		t.Fatalf("expected selected index %d, got %d", want, got)
	}

	if got, want := model.SelectedChoice(), "on_track"; got != want {
		t.Fatalf("expected selected choice %q, got %q", want, got)
	}
}

func TestArrowNavigation(t *testing.T) {
	model := NewModel([]string{"on_track", "off_track", "deep_in_the_void", "snooze"})
	updated, _ := model.Update(tea.KeyMsg{Type: tea.KeyDown})
	next, ok := updated.(Model)
	if !ok {
		t.Fatal("expected updated model type")
	}

	if got, want := next.SelectedChoice(), "off_track"; got != want {
		t.Fatalf("expected selected choice %q, got %q", want, got)
	}
}
