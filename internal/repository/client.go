package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"smartparking/internal/models"
	"smartparking/pkg/logger"
	"strings"
)

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *clientRepository {
	return &clientRepository{db: db}
}

func (repo *clientRepository) GetByID(id int64) (result models.Client, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("clientRepository.GetByID failed", zap.Error(err), zap.Int64("id", id))
		}
	}()

	err = repo.db.Model(&models.Client{}).Where("id = ?", id).First(&result).Error
	return
}

func (repo *clientRepository) GetByPhone(phone string) (result models.Client, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("clientRepository.GetByPhone failed", zap.Error(err), zap.String("phone", phone))
		}
	}()

	err = repo.db.Model(&models.Client{}).Where("phone = ?", phone).First(&result).Error
	return
}

func (repo *clientRepository) GetByEmail(email string) (result models.Client, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("clientRepository.GetByPhone GetByEmail", zap.Error(err), zap.String("email", email))
		}
	}()

	err = repo.db.Model(&models.Client{}).Where("LOWER(email) = LOWER(?)", email).First(&result).Error
	return
}

func (repo *clientRepository) GetAll() (result []models.Client, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("clientRepository.GetAll failed", zap.Error(err))
		}
	}()

	err = repo.db.Model(&models.Client{}).Find(&result).Error
	return
}

func (repo *clientRepository) Create(model models.Client) (result models.Client, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("clientRepository.Create", zap.Error(err), zap.Any("model", model))
		}
	}()

	model.Email = strings.ToLower(model.Email)
	err = repo.db.Model(&models.Client{}).Create(&model).Error
	if err != nil {
		return
	}

	return repo.GetByID(model.ID)
}

func (repo *clientRepository) Update(id int64, model models.Client) (result models.Client, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("clientRepository.Update failed", zap.Error(err), zap.Int64("id", id), zap.Any("model", model))
		}
	}()

	model.Email = strings.ToLower(model.Email)
	err = repo.db.Model(&models.Client{}).Where("id = ?", id).Updates(&model).Error
	if err != nil {
		return
	}

	return repo.GetByID(id)
}

func (repo *clientRepository) UpdateByEmail(email string, model models.Client) (result models.Client, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("clientRepository.UpdateByEmail failed", zap.Error(err), zap.String("email", email), zap.Any("model", model))
		}
	}()

	err = repo.db.Model(&models.Client{}).Where("LOWER(email) = LOWER(?)", email).Omit("email").Updates(&model).Error
	if err != nil {
		return
	}

	return repo.GetByEmail(email)
}

func (repo *clientRepository) Delete(id int64) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("clientRepository.Delete failed", zap.Error(err), zap.Int64("id", id))
		}
	}()

	err = repo.db.Model(&models.Client{}).Where("id = ?", id).Delete(&models.Client{}).Error
	return
}
