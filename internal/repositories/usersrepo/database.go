package usersrepo

import (
	"database/sql"

	"github.com/wesleynepo/auth-service-go/internal/core/domain"
)

type relational struct {
   db *sql.DB
}

func NewRelational(db *sql.DB) *relational {
	return &relational{db: db}
}

func (r *relational) FindByMail(email string) (domain.User, error) {
    statement := "SELECT * FROM users WHERE email = $1"
    user := domain.User{}

    row := r.db.QueryRow(statement, email)

    err := row.Scan(&user.ID, &user.Email, &user.Hash)

    if err != nil {
        return user, err
    }

    return user, nil
}

func (r *relational) Save(user domain.User) error {
    statement := "INSERT INTO users (email, hash) VALUES ($1, $2)"

    _, err := r.db.Exec(statement, user.Email, user.Hash)

    return err
}

