package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"math/rand"
	"time"
)

type TokenManager interface {
	NewJWT(claims *Claims) (string, error)
	Parse(accessToken string) (*Claims, error)
	NewRefreshToken() (string, error)
}

type manager struct {
	secret string
}

func NewManager(secret string) *manager {
	return &manager{secret: secret}
}

func (m *manager) NewJWT(claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.secret))
}

func (m *manager) Parse(accessToken string) (*Claims, error) {
	claims := Claims{}
	_, err := jwt.ParseWithClaims(accessToken, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Header["alg"])
		}
		return []byte(m.secret), nil
	})
	if err != nil {
		return nil, err
	}

	err = claims.Valid()
	if err != nil {
		return nil, err
	}

	return &claims, nil
}

func (m *manager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}
