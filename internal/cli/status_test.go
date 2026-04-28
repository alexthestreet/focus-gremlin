package cli

import "testing"

func TestStatusCommandExists(t *testing.T) {
	cmd := NewRootCommand()
	if child := findSubcommand(cmd, "status"); child == nil {
		t.Fatal("expected status subcommand")
	}
}
