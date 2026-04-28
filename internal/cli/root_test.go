package cli

import "testing"

func TestRootCommandExists(t *testing.T) {
	cmd := NewRootCommand()
	if cmd == nil {
		t.Fatal("expected root command")
	}

	if got, want := cmd.Use, "focus-gremlin"; got != want {
		t.Fatalf("expected use %q, got %q", want, got)
	}

	expected := []string{"daemon", "prompt", "config"}
	for _, use := range expected {
		if child := findSubcommand(cmd, use); child == nil {
			t.Fatalf("expected subcommand %q", use)
		}
	}
}

func findSubcommand(cmd *Command, use string) *Command {
	for _, child := range cmd.Commands() {
		if child.Use == use {
			return child
		}
	}

	return nil
}
