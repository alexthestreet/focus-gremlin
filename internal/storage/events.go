package storage

import "time"

type Event struct {
	ID         int64
	Status     string
	RecordedAt time.Time
}
