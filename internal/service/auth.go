package service

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"math/rand"
	"smartparking/config"
	"smartparking/internal/apiError"
	"smartparking/internal/interfaces/manager"
	"smartparking/internal/models"
	toolsEmail "smartparking/pkg/email"
	"smartparking/pkg/hash"
	"smartparking/pkg/jwt"
	"smartparking/pkg/logger"
	"time"
)

type authService struct {
	m            manager.Manager
	tokenManager jwt.TokenManager
	hashManager  hash.Manager
	emailManager toolsEmail.EmailManager
}

func NewAuthService(m manager.Manager, tokenManager jwt.TokenManager, hashManager hash.Manager, emailManager toolsEmail.EmailManager) *authService {
	return &authService{
		m:            m,
		tokenManager: tokenManager,
		hashManager:  hashManager,
		emailManager: emailManager,
	}
}

func (s *authService) SignUp(client models.Client) (result models.Client, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("authService.Register failed", zap.Error(err), zap.Any("client", client))
		}
	}()

	err = s.checkAlreadyExistenceByEmailAndPhone(client)
	if err != nil {
		return
	}

	hashedPassword, err := s.hashManager.HashPassword(client.Password)
	if err != nil {
		return
	}

	client.Password = hashedPassword
	return s.m.Service().Client().Create(client)
}

func (s *authService) SignIn(client models.Client) (tokens models.Tokens, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("authService.Login failed", zap.Error(err), zap.Any("client", client))
		}
	}()

	result, err := s.m.Service().Client().GetByEmail(client.Email)
	if err != nil {
		return
	}

	if !s.hashManager.CheckPasswordHash(client.Password, result.Password) {
		err = apiError.Throw(apiError.InvalidPassword)
		return
	}

	return s.createSession(result.ID)
}

func (s *authService) ValidateToken(token string) (clientID int64, err error) {
	return s.tokenManager.Parse(token)
}

func (s *authService) RefreshTokens(refreshToken string) (tokens models.Tokens, err error) {
	session, err := s.m.Repository().Session().GetByRefreshToken(refreshToken)
	if err != nil {
		return
	}

	if session.IsExpired() {
		err = apiError.Throw(apiError.TokenExpired)
		return
	}

	newAccessToken, err := s.tokenManager.NewJWT(session.ClientID, config.GlobalConfig.JWT.AccessTokenTTL)
	if err != nil {
		return
	}

	tokens = models.Tokens{
		AccessToken:  newAccessToken,
		RefreshToken: session.RefreshToken,
	}
	return
}

func (s *authService) createSession(clientID int64) (models.Tokens, error) {
	accessToken, err := s.tokenManager.NewJWT(clientID, config.GlobalConfig.JWT.AccessTokenTTL)
	if err != nil {
		return models.Tokens{}, err
	}

	refreshToken, err := s.tokenManager.NewRefreshToken()
	if err != nil {
		return models.Tokens{}, err
	}

	session := models.Session{
		ClientID:     clientID,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(config.GlobalConfig.JWT.RefreshTokenTTL),
	}

	_, err = s.m.Repository().Session().Create(session)
	if err != nil {
		return models.Tokens{}, err
	}

	return models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) GenerateEmailVerificationAndSendToClient(email, password string) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("authService.UpdatePasswordByEmailVerification failed", zap.Error(err), zap.String("email", email))
		}
	}()

	_, err = s.m.Service().Client().GetByEmail(email)
	if err != nil {
		return
	}

	code := s.generateVerificationCode()
	err = s.emailManager.Send(email, code)
	if err != nil {
		return err
	}

	hashedPassword, err := s.hashManager.HashPassword(password)
	if err != nil {
		return err
	}

	emailVerification := models.EmailVerification{
		Email:         email,
		Password:      hashedPassword,
		Code:          code,
		GeneratedTime: time.Now(),
	}
	_, err = s.m.Repository().EmailVerification().Create(emailVerification)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) CheckEmailVerificationAndUpdatePassword(email, code string) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("authService.CheckEmailVerification failed", zap.Error(err), zap.String("email", email), zap.String("code", code))
		}
	}()

	result, err := s.m.Repository().EmailVerification().GetByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = apiError.Throw(apiError.ClientNotFound)
		}
		return
	}

	err = result.IsValid(code)
	if err != nil {
		return
	}

	err = s.m.Repository().EmailVerification().SetCheckedByID(result.ID, true)
	if err != nil {
		return
	}

	_, err = s.m.Repository().Client().UpdateByEmail(email, models.Client{
		Password: result.Password,
	})
	if err != nil {
		return
	}

	return nil
}

func (s *authService) generateVerificationCode() string {
	code := 1000 + rand.Intn(900)
	return fmt.Sprintf("%d", code)
}

func (s *authService) checkAlreadyExistenceByEmailAndPhone(client models.Client) error {
	result, err := s.m.Repository().Client().GetByEmail(client.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if result.ID != 0 {
		return apiError.Throw(apiError.ClientWithEmailExists)
	}

	result, err = s.m.Repository().Client().GetByPhone(client.Phone)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if result.ID != 0 {
		return apiError.Throw(apiError.ClientWithPhoneExists)
	}

	return nil
}
