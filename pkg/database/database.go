package database

import (
	"database/sql"
	"log"
)

type Storage struct {
   db *sql.DB
}

func New(dsn string) *Storage {
    db, err := sql.Open("postgres", dsn)

    if err != nil {
        log.Fatal(err)
        return nil
    }

    return &Storage{db: db}
}


func (s Storage) Close() {
    s.db.Close()
}

func (s Storage) Get() (*sql.DB) {
    return s.db
}
