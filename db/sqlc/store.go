package db

import (
	"database/sql"
)

// extends the sqlc queries objects to perform transactions
type Store struct {
	*Queries
	db *sql.DB // pass a db object
}

// Newstore function creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
		// New function generated by sqlc
		Queries: New(db),
	}
}
