package usersrv

import (
	"errors"

	"github.com/wesleynepo/auth-service-go/internal/core/domain"
	"github.com/wesleynepo/auth-service-go/internal/core/ports"
	"github.com/wesleynepo/auth-service-go/pkg/hash"
)

type service struct {
    usersRepository ports.UsersRepository
    hasher hash.Hasher
 }

const TOKEN_DURATION = 4
const REFRESH_TOKEN_DURATION = 12

func New(usersRepository ports.UsersRepository, hasher hash.Hasher) *service {
    return &service{usersRepository: usersRepository, hasher: hasher}
}

func (s *service) CheckCredentials(email, password string) (uint, error) {
    user, err := s.usersRepository.FindByMail(email)

    if err != nil {
        return 0, err
    }

    valid := s.hasher.PasswordMatches(password, user.Hash)

    if !valid {
        return 0, errors.New("invalid password")
    }

    return user.ID, nil
}

func (s *service) Create(email, password, confirmPassword string) (error) {
    if password != confirmPassword {
        return errors.New("Password and confirmPassword doesn't match")
    }

    hash, err := s.hasher.HashPassword(password)

    if err != nil {
        return err
    }

    user := domain.User{Email: email, Hash: hash}

    err = s.usersRepository.Save(user)

    return err
}
