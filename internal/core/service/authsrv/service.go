package authsrv

import (
	"time"

	"github.com/wesleynepo/auth-service-go/internal/core/domain"
	"github.com/wesleynepo/auth-service-go/pkg/jwt"
)

type service struct {
    jwt jwt.JWTGen
}

const TOKEN_DURATION = 4
const REFRESH_TOKEN_DURATION = 12

func New(jwt jwt.JWTGen) *service {
    return &service{jwt: jwt}
}

func tokenTime() (time.Time) {
    return time.Now().Add(time.Hour * TOKEN_DURATION)
}

func refreshTokenTime() (time.Time) {
    return time.Now().Add(time.Hour * REFRESH_TOKEN_DURATION)
}

func (s *service) Login(email, password string) (domain.Auth, error) {

    //TODO: user service validation
    refreshTokenDuration := time.Now().Add(time.Hour * 12)
    userId := 123.0

    token, err := s.jwt.CreateToken(userId, tokenTime())

    if (err != nil) {
        return domain.Auth{}, err
    }

    refresh, err := s.jwt.CreateToken(userId, refreshTokenDuration)

    if (err != nil) {
        return domain.Auth{}, err
    }

    return domain.Auth{Token: token, RefreshToken: refresh}, nil
}

func (s *service) Refresh(refresh string) (domain.Auth, error) {
    id, err := s.jwt.CheckToken(refresh)

    if err != nil {
        return domain.Auth{}, err
    }

    token, err := s.jwt.CreateToken(id, tokenTime())

    if err != nil {
        return domain.Auth{}, err
    }

    return domain.Auth{Token: token, RefreshToken: refresh}, nil
}
