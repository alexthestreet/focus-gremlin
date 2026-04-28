package daemon

import (
	"path/filepath"
	"testing"
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
