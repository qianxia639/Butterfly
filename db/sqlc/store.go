package db

import (
	"github.com/jmoiron/sqlx"
)

type Store struct {
	*Queries
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
