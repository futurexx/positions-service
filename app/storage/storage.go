package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Storage is layout between handlers and database
type Storage struct {
	config *Config
	db     *sql.DB
}

// New init Store with config
func New(config *Config) *Storage {
	return &Storage{config: config}
}

// Open new connection to database
func (s *Storage) Open() error {
	db, err := sql.Open("sqlite3", s.config.DatabaseFile)
	if err != nil {
		return err
	}

	s.db = db
	return nil
}

// Close database connection
func (s *Storage) Close() error {
	if err := s.db.Close(); err != nil {
		return err
	}

	return nil
}
