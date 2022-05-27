package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	DB *sqlx.DB
}

func NewDB(dsn string) (*DB, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &DB{DB: db}, nil
}
