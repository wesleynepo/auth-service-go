package authsrv

import (
	"time"

	"github.com/wesleynepo/auth-service-go/internal/core/domain"
	"github.com/wesleynepo/auth-service-go/pkg/jwt"
)

type service struct {
    jwt jwt.JWTGen
}

func New(jwt jwt.JWTGen) *service {
    return &service{jwt: jwt}
}


func (s *service) Login(email, password string) (domain.Auth, error) {

    //TODO: user service validation
    tokenDuration := time.Now().Add(time.Hour * 4)
    refreshTokenDuration := time.Now().Add(time.Hour * 12)
    userId := 123.0

    token, err := s.jwt.CreateToken(userId, tokenDuration)

    if (err != nil) {
        return domain.Auth{}, err
    }

    refresh, err := s.jwt.CreateToken(userId, refreshTokenDuration)

    if (err != nil) {
        return domain.Auth{}, err
    }

    return domain.Auth{Token: token, RefreshToken: refresh}, nil
}
