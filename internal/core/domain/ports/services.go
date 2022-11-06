package ports

import "github.com/wesleynepo/auth-service-go/internal/core/domain"

type AuthService interface {
    Refresh(refresh string) (domain.Auth, error)
    Login(login, password string) (domain.Auth, error)
}
