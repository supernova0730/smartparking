package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"math/rand"
	"smartparking/pkg/tools"
	"time"
)

type TokenManager interface {
	NewJWT(clientID int64, ttl time.Duration) (string, error)
	Parse(accessToken string) (int64, error)
	NewRefreshToken() (string, error)
}

type Manager struct {
	secret string
}

func NewManager(secret string) *Manager {
	return &Manager{secret: secret}
}

func (m *Manager) NewJWT(clientID int64, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(ttl).Unix(),
		Subject:   fmt.Sprintf("%d", clientID),
	})

	return token.SignedString([]byte(m.secret))
}

func (m *Manager) Parse(accessToken string) (int64, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", token.Header["alg"])
		}
		return []byte(m.secret), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("error get user claims from token")
	}

	clientIDstr := claims["sub"].(string)
	clientID := tools.StringToInt64(clientIDstr)

	return clientID, nil
}

func (m *Manager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	_, err := r.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}
