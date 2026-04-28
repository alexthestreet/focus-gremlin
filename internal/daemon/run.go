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
	if options.PollInterval <= 0 {
		options.PollInterval = time.Second
	}

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

		action, err := scheduler.NextAction(now, options.Scheduler, state)
		if err != nil {
			return err
		}

		if action.SkipPrompt {
			if err := waitFor(ctx, options.PollInterval); err != nil {
				return nil
			}
			continue
		}

		if wait := time.Until(action.When); wait > 0 {
			if err := waitFor(ctx, wait); err != nil {
				return nil
			}
			continue
		}

		if err := markPromptActive(options.StatePath); err != nil {
			return err
		}

		cmd, err := options.Launch(options.TerminalCommand, options.AppCommand)
		if err != nil {
			clearPromptActive(options.StatePath)
			return fmt.Errorf("launch prompt: %w", err)
		}

		go waitForPromptExit(cmd, options.StatePath)
		if err := waitFor(ctx, options.PollInterval); err != nil {
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
