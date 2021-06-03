package database

import "database/sql"

type PostDB interface {
	Open() error
	Close() error
	emailExists(email string) bool
}

type DB struct {
	db *sql.DB
}

func (d *DB) Open() error {
	return nil
}

func (d *DB) Close() error {
	return nil
}
