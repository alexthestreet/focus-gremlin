package checkin

import "github.com/alexthestreet/focus-gremlin/internal/storage"

func SubmitResponse(store *storage.Store, result Result) error {
	return store.InsertEvent(storage.Event{
		Status:     result.Status,
		RecordedAt: result.RecordedAt,
	})
}
