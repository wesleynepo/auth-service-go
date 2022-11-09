package hash

import "golang.org/x/crypto/bcrypt"

const cost = 10

type Hasher interface {
    HashPassword(password string) (string, error)
    PasswordMatches(password,hash string) (bool)
}

type crypt struct {}

func New() Hasher {
    return &crypt{}
}

func (b crypt) HashPassword(password string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
    if err != nil {
        return "", nil
    }
    return string(hash), nil
}

func (b crypt) PasswordMatches(password, hash string) (bool) {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
