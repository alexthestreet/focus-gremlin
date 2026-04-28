package daemon

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/alexthestreet/focus-gremlin/internal/launcher"
	runtimestate "github.com/alexthestreet/focus-gremlin/internal/runtime"
	"github.com/alexthestreet/focus-gremlin/internal/scheduler"
)

type Options struct {
	LockPath        string
	StatePath       string
	Scheduler       scheduler.Options
	TerminalCommand []string
	AppCommand      []string
	Now             func() time.Time
	Launch          func([]string, []string) (*exec.Cmd, error)
	Sleep           func(context.Context, time.Duration) error
	PollInterval    time.Duration
}

func Run(ctx context.Context, options Options) error {
	lock, err := acquireDaemonLock(options.LockPath)
	if err != nil {
		return err
	}
	defer lock.Close()

	if options.Now == nil {
		options.Now = time.Now
	}
	if options.Launch == nil {
		options.Launch = launcher.Start
	}
	if options.Sleep == nil {
		options.Sleep = waitFor
	}
	if options.PollInterval <= 0 {
		options.PollInterval = time.Second
	}

	var scheduledAt time.Time

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		now := options.Now().UTC()
		state, err := runtimestate.LoadState(options.StatePath)
		if err != nil {
			return err
		}

		if state.PromptActive {
			if err := options.Sleep(ctx, options.PollInterval); err != nil {
				return nil
			}
			continue
		}

		if !state.SnoozedUntil.IsZero() && state.SnoozedUntil.After(now) {
			scheduledAt = state.SnoozedUntil.UTC()
		} else if scheduledAt.IsZero() {
			scheduledAt, err = scheduler.NextCheckIn(now, options.Scheduler)
			if err != nil {
				return err
			}
		}

		if wait := scheduledAt.Sub(now); wait > 0 {
			if err := options.Sleep(ctx, wait); err != nil {
				return nil
			}
			continue
		}

		if err := markPromptActive(options.StatePath); err != nil {
			return err
		}
		scheduledAt = time.Time{}

		cmd, err := options.Launch(options.TerminalCommand, options.AppCommand)
		if err != nil {
			clearPromptActive(options.StatePath)
			return fmt.Errorf("launch prompt: %w", err)
		}

		go waitForPromptExit(cmd, options.StatePath)
		if err := options.Sleep(ctx, options.PollInterval); err != nil {
			return nil
		}
	}
}

func acquireDaemonLock(path string) (*runtimestate.Lock, error) {
	return runtimestate.AcquireLock(path)
}

func waitFor(ctx context.Context, duration time.Duration) error {
	timer := time.NewTimer(duration)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func markPromptActive(path string) error {
	state, err := runtimestate.LoadState(path)
	if err != nil {
		return err
	}

	state.PromptActive = true
	return runtimestate.SaveState(path, state)
}

func clearPromptActive(path string) error {
	state, err := runtimestate.LoadState(path)
	if err != nil {
		return err
	}

	state.PromptActive = false
	return runtimestate.SaveState(path, state)
}

func waitForPromptExit(cmd *exec.Cmd, statePath string) {
	if cmd == nil {
		clearPromptActive(statePath)
		return
	}

	_ = cmd.Wait()
	_ = clearPromptActive(statePath)
}
