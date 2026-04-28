package scheduler

import "time"

type Options struct {
	Interval    time.Duration
	ActiveStart string
	ActiveEnd   string
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
