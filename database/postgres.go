package database

import (
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func Connect(addr string) (*sql.DB, error) {
	db, err := sql.Open("pgx", addr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
