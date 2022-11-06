package jwt

import (
	"errors"
	"time"

	"github.com/pascaldekloe/jwt"
)

type JWTGen interface {
    CreateToken(userId uint, expiration time.Time) (string, error)
    CheckToken(token string) (uint, error)
}

type jwtgen struct {}

func New() JWTGen {
    return &jwtgen{}
}

func (j jwtgen) CreateToken(userId uint, expiration time.Time) (string, error) {
    var claims jwt.Claims

    claims.Set = map[string]interface{}{"id": userId}
    claims.Issued  = jwt.NewNumericTime(time.Now().Round(time.Second))
    claims.Expires = jwt.NewNumericTime(expiration)

    token, err := claims.HMACSign("HS256", []byte("test"))

    if (err != nil) {
        return "", errors.New("Couldn't sign the claims")
    }

    return string(token), nil
}

func (j jwtgen) CheckToken(token string) (uint, error) {
    claims, err := jwt.HMACCheck([]byte(token), []byte("test"))

    if err != nil {
        return 0, errors.New("Invalid token")
    }

    if (!claims.Valid(time.Time{})) {
        return 0, errors.New("Token expired")
    }

    return claims.Set["id"].(uint), nil
}
