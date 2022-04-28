package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type Manager interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type manager struct{}

func NewManager() *manager {
	return &manager{}
}

func (h *manager) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (h *manager) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
