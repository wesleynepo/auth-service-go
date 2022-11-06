package usersrv

import (
	"errors"

	"github.com/wesleynepo/auth-service-go/internal/core/domain"
	"github.com/wesleynepo/auth-service-go/internal/core/ports"
)

type service struct {
    usersRepository ports.UsersRepository
}

const TOKEN_DURATION = 4
const REFRESH_TOKEN_DURATION = 12

func New(usersRepository ports.UsersRepository) *service {
    return &service{usersRepository: usersRepository}
}

func (s *service) CheckCredentials(email, password string) (uint, error) {
    user, err := s.usersRepository.FindByMail(email)

    if err != nil {
        return 0, err
    }

    // Logic with hash and salt for password
    valid := true

    if !valid {
        return 0, errors.New("invalid password")
    }

    return user.ID, nil
}

func (s *service) Create(email, password, confirmPassword string) (error) {
    if password != confirmPassword {
        return errors.New("Password and confirmPassword doesn't match")
    }

    // HASH AND SALT PASSWORD LOGIC
    hash := "AA"
    salt := "BB"

    user := domain.User{Email: email, Hash: hash, Salt: salt}

    err := s.usersRepository.Save(user)

    return err
}
