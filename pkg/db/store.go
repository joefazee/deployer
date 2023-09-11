package db

import (
	"database/sql"

	"github.com/joefazee/autodeploy/pkg/domain"
)

type SQLStore struct {
	db *sql.DB
}

func NewDBStore(db *sql.DB) domain.Store {
	return &SQLStore{
		db: db,
	}
}
