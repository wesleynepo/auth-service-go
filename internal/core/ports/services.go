package ports

import "github.com/wesleynepo/auth-service-go/internal/core/domain"

type AuthService interface {
    Refresh(email, password string) (domain.Auth, error)
    Login(user domain.User) (domain.Auth, error)
}
