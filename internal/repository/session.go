package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"smartparking/internal/models"
	"smartparking/pkg/logger"
)

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) *sessionRepository {
	return &sessionRepository{db: db}
}

func (repo *sessionRepository) GetByID(id int64) (result models.Session, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("sessionRepository.GetByID failed", zap.Error(err), zap.Int64("id", id))
		}
	}()

	err = repo.db.
		Model(&models.Session{}).
		Where("id = ?", id).
		First(&result).
		Error
	return
}

func (repo *sessionRepository) GetByRefreshToken(refreshToken string) (result models.Session, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("sessionRepository.GetByRefreshToken failed", zap.Error(err), zap.String("refreshToken", refreshToken))
		}
	}()

	err = repo.db.
		Model(&models.Session{}).
		Where("refresh_token = ?", refreshToken).
		Order("id desc").
		First(&result).
		Error
	return
}

func (repo *sessionRepository) Create(model models.Session) (result models.Session, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("sessionRepository.Create failed", zap.Error(err), zap.Any("model", model))
		}
	}()

	err = repo.db.Model(&models.Session{}).Create(&model).Error
	if err != nil {
		return
	}

	return repo.GetByID(model.ID)
}

func (repo *sessionRepository) DeleteByID(id int64) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("sessionRepository.DeleteByID failed", zap.Error(err), zap.Int64("id", id))
		}
	}()

	err = repo.db.Model(&models.Session{}).Where("id = ?", id).Delete(&models.Session{}).Error
	return
}
