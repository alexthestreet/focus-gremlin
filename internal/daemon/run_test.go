package daemon

import (
	"context"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/alexthestreet/focus-gremlin/internal/scheduler"
)

func TestSingleInstanceLock(t *testing.T) {
	path := filepath.Join(t.TempDir(), "daemon.lock")
	first, err := acquireDaemonLock(path)
	if err != nil {
		t.Fatalf("acquire first lock: %v", err)
	}
	defer first.Close()

	second, err := acquireDaemonLock(path)
	if err == nil {
		second.Close()
		t.Fatal("expected second lock acquisition to fail")
	}
}

func TestRunLaunchesPromptWhenCheckInIsDue(t *testing.T) {
	root := t.TempDir()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nowValues := []time.Time{
		time.Date(2026, time.April, 28, 10, 0, 0, 0, time.UTC),
		time.Date(2026, time.April, 28, 11, 0, 0, 0, time.UTC),
	}
	index := 0
	launches := 0
	err := Run(ctx, Options{
		LockPath:  filepath.Join(root, "daemon.lock"),
		StatePath: filepath.Join(root, "state.json"),
		Scheduler: scheduler.Options{
			Interval:    time.Hour,
			ActiveStart: "09:00",
			ActiveEnd:   "17:00",
		},
		Now: func() time.Time {
			if index >= len(nowValues) {
				return nowValues[len(nowValues)-1]
			}
			value := nowValues[index]
			index++
			return value
		},
		Sleep: func(ctx context.Context, duration time.Duration) error {
			return nil
		},
		Launch: func(terminalCommand []string, appCommand []string) (*exec.Cmd, error) {
			launches++
			cancel()
			return nil, nil
		},
	})
	if err != nil {
		t.Fatalf("run daemon: %v", err)
	}

	if launches != 1 {
		t.Fatalf("expected 1 prompt launch, got %d", launches)
	}
}
