package launcher

import "testing"

func TestBuildCommand(t *testing.T) {
	command, err := BuildCommand([]string{"kitty", "--hold", "-e"}, []string{"focus-gremlin", "prompt"})
	if err != nil {
		t.Fatalf("build command: %v", err)
	}

	if got, want := command.Path, "kitty"; got != want {
		t.Fatalf("expected path %q, got %q", want, got)
	}

	wantArgs := []string{"kitty", "--hold", "-e", "focus-gremlin", "prompt"}
	if len(command.Args) != len(wantArgs) {
		t.Fatalf("expected %d args, got %d", len(wantArgs), len(command.Args))
	}

	for i := range wantArgs {
		if command.Args[i] != wantArgs[i] {
			t.Fatalf("expected arg %d to be %q, got %q", i, wantArgs[i], command.Args[i])
		}
	}
}
