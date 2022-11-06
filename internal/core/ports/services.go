package ports

import "github.com/wesleynepo/auth-service-go/internal/core/domain"

type AuthService interface {
    Refresh(refresh string) (domain.Auth, error)
    Login(email, password string) (domain.Auth, error)
}

type UserService interface {
    CheckCredentials(email, password string) (float64, error)
}
