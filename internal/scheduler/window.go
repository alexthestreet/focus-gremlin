package scheduler

import (
	"fmt"
	"time"
)

func ParseClock(value string) (int, int, error) {
	parsed, err := time.Parse("15:04", value)
	if err != nil {
		return 0, 0, fmt.Errorf("parse clock value %q: %w", value, err)
	}

	return parsed.Hour(), parsed.Minute(), nil
}

func AtClock(base time.Time, value string) (time.Time, error) {
	hour, minute, err := ParseClock(value)
	if err != nil {
		return time.Time{}, err
	}

	return time.Date(base.Year(), base.Month(), base.Day(), hour, minute, 0, 0, base.Location()), nil
}
