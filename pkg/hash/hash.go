package hash

import (
	"crypto/sha256"
	"fmt"
)

type HashManager interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type SHA256Hasher struct {
	salt string
}

func NewManager(salt string) *SHA256Hasher {
	return &SHA256Hasher{salt: salt}
}

func (h *SHA256Hasher) HashPassword(password string) (string, error) {
	hash := sha256.New()

	_, err := hash.Write([]byte(password))
	if err != nil {
		return "", err
	}

	_, err = hash.Write([]byte(h.salt))
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%x", hash.Sum(nil))
	return result, nil
}

func (h *SHA256Hasher) CheckPasswordHash(password, hash string) bool {
	hashedPassword, err := h.HashPassword(password)
	if err != nil {
		return false
	}

	return hashedPassword == hash
}
