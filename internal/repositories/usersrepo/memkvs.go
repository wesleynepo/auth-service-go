package usersrepo

import (
	"encoding/json"
	"errors"

	"github.com/wesleynepo/auth-service-go/internal/core/domain"
)

type memkvs struct {
	kvs map[string][]byte
}

func NewMemKVS() *memkvs {
	return &memkvs{kvs: map[string][]byte{}}
}

func (repo *memkvs) FindByMail(email string) (domain.User, error) {
	if value, ok := repo.kvs[email]; ok {
		user := domain.User{}
		err := json.Unmarshal(value, &user)

		if err != nil {
			return domain.User{}, err
		}

		return user, nil
	}

	return domain.User{}, errors.New("user not found in kvs")
}

func (repo *memkvs) Save(user domain.User) error {
	bytes, err := json.Marshal(user)

	if err != nil {
		return errors.New("user fails at marshal into json string")
	}

	repo.kvs[user.Email] = bytes

	return nil
}

