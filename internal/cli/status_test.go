package cli

import (
	"path/filepath"
	"testing"
)

func TestStatusCommandExists(t *testing.T) {
	cmd := NewRootCommand()
	if child := findSubcommand(cmd, "status"); child == nil {
		t.Fatal("expected status subcommand")
	}
}

func TestPromptAppCommandUsesExecutablePath(t *testing.T) {
	command := promptAppCommand(filepath.Join("/tmp", "focus-gremlin"))
	want := []string{filepath.Join("/tmp", "focus-gremlin"), "prompt"}
	for i := range want {
		if command[i] != want[i] {
			t.Fatalf("expected arg %d to be %q, got %q", i, want[i], command[i])
		}
	}
}
