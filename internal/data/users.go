package data

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)
var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

type User struct {
	ID int64 `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password password `json:"password"`
	Activated bool `json:"activated"`
	Version int `json:"version"`
}

type password struct {
	plaintext *string
	hash []byte
}

func (p *password) Set(plainTextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainTextPassword), 12)
	if err != nil {
		return err
	}

	p.plaintext = &plainTextPassword
	p.hash = hash
	return nil
}

func (p *password) Matches(plainTextPassword string) (bool, error){
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plainTextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}