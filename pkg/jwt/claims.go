package jwt

import (
	"github.com/golang-jwt/jwt"
	"smartparking/pkg/tools"
	"strconv"
	"time"
)

type Claims struct {
	jwt.StandardClaims
}

func (c *Claims) GetSubject() int64 {
	return tools.StringToInt64(c.Subject)
}

func (c *Claims) SetSubject(subject int64) {
	c.Subject = strconv.FormatInt(subject, 10)
}

func (c *Claims) SetExpiresAt(ttl time.Duration) {
	c.ExpiresAt = time.Now().Add(ttl).Unix()
}
