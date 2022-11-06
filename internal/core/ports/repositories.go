package ports

import "github.com/wesleynepo/auth-service-go/internal/core/domain"

type UsersRepository interface {
    FindByMail(email string) (domain.User, error)
}
