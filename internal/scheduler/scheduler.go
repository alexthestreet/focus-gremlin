package scheduler

import (
	"time"

	runtimestate "github.com/alexthestreet/focus-gremlin/internal/runtime"
)

type Options struct {
	Interval    time.Duration
	ActiveStart string
	ActiveEnd   string
}

type Action struct {
	When       time.Time
	SkipPrompt bool
}

func NextCheckIn(now time.Time, options Options) (time.Time, error) {
	next := now.Add(options.Interval)
	start, err := AtClock(next, options.ActiveStart)
	if err != nil {
		return time.Time{}, err
	}

	end, err := AtClock(next, options.ActiveEnd)
	if err != nil {
		return time.Time{}, err
	}

	if next.Before(start) {
		return start, nil
	}

	if next.After(end) {
		tomorrow := next.AddDate(0, 0, 1)
		return AtClock(tomorrow, options.ActiveStart)
	}

	return next, nil
}

func NextAction(now time.Time, options Options, state runtimestate.State) (Action, error) {
	if state.PromptActive {
		return Action{When: now, SkipPrompt: true}, nil
	}

	if !state.SnoozedUntil.IsZero() && state.SnoozedUntil.After(now) {
		return Action{When: state.SnoozedUntil}, nil
	}

	next, err := NextCheckIn(now, options)
	if err != nil {
		return Action{}, err
	}

	return Action{When: next}, nil
}
