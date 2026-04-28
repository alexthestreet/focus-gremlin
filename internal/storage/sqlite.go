package storage

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

type Store struct {
	db *sql.DB
}

func Open(path string) (*Store, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, fmt.Errorf("create storage dir for %q: %w", path, err)
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("open sqlite database %q: %w", path, err)
	}

	store := &Store{db: db}
	if err := store.initSchema(); err != nil {
		db.Close()
		return nil, err
	}

	return store, nil
}

func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}

	return s.db.Close()
}

func (s *Store) InsertEvent(event Event) error {
	_, err := s.db.Exec(`
		INSERT INTO events (status, recorded_at)
		VALUES (?, ?)
	`, event.Status, event.RecordedAt.UTC())
	if err != nil {
		return fmt.Errorf("insert event: %w", err)
	}

	return nil
}

func (s *Store) RecentEvents(limit int) ([]Event, error) {
	rows, err := s.db.Query(`
		SELECT id, status, recorded_at
		FROM events
		ORDER BY recorded_at DESC, id DESC
		LIMIT ?
	`, limit)
	if err != nil {
		return nil, fmt.Errorf("query recent events: %w", err)
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.Status, &event.RecordedAt); err != nil {
			return nil, fmt.Errorf("scan event: %w", err)
		}
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate recent events: %w", err)
	}

	return events, nil
}

func (s *Store) initSchema() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			status TEXT NOT NULL,
			recorded_at TIMESTAMP NOT NULL
		)
	`)
	if err != nil {
		return fmt.Errorf("create events schema: %w", err)
	}

	return nil
}
