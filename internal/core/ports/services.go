package ports

type AuthService interface {
    Refresh(token string, refresh string) (token, refresh string, error)
    Login(user domain.User) (token, refresh string, error)
}
