package service

import (
	"smartparking/internal/models"
	"smartparking/pkg/jwt"
)

type Auth interface {
	SignUp(model models.Client) (result models.Client, err error)
	SignIn(model models.Client) (tokens models.Tokens, err error)
	ValidateToken(token string) (claims *jwt.Claims, err error)
	RefreshTokens(refreshToken string) (tokens models.Tokens, err error)
	GenerateEmailVerificationAndSendToClient(email, password string) (err error)
	CheckEmailVerificationAndUpdatePassword(email, code string) (err error)
}
