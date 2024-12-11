package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Repo struct {
	db *sql.DB
}

func New() (*Repo, error) {
	// TODO: use env variable for db name
	db, err := sql.Open("sqlite3", "pdty.db")
	if err != nil {
		return nil, err
	}
	return &Repo{
		db: db,
	}, nil
}
