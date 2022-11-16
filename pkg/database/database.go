package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Storage struct {
   db *sql.DB
}


const DSN_TEMPLATE = "postgres://%s:%s@localhost:5432/postgres?sslmode=disable&application_name=%s"

func New() *Storage {
    user := os.Getenv("POSTGRES_USER")
    password := os.Getenv("POSTGRES_PASSWORD")
    applicationName := os.Getenv("POSTGRES_APPLICATION_NAME")
    db, err := sql.Open("postgres", fmt.Sprintf(DSN_TEMPLATE, user, password, applicationName))

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
