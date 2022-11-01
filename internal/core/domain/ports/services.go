package ports

import "github.com/wesleynepo/auth-service-go/internal/core/domain"

type AuthService interface {
    Refresh(domain.Auth) (domain.Auth, error)
    Login(domain.User) (domain.Auth, error)
}
