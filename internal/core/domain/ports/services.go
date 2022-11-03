package ports

import "github.com/wesleynepo/auth-service-go/internal/core/domain"

type AuthService interface {
//    Refresh(domain.Auth) (domain.Auth, error)
    Login(login, password string) (domain.Auth, error)
}
