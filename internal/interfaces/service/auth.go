package service

import (
	"smartparking/internal/models"
)

type Auth interface {
	SignUp(model models.Client) (result models.Client, err error)
	SignIn(model models.Client) (tokens models.Tokens, err error)
	ValidateToken(token string) (clientID int64, err error)
	RefreshTokens(refreshToken string) (tokens models.Tokens, err error)
	GenerateEmailVerificationAndSendToClient(email, password string) (err error)
	CheckEmailVerificationAndUpdatePassword(email, code string) (err error)
}
