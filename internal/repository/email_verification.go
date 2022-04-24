package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"smartparking/internal/models"
	"smartparking/pkg/logger"
)

type emailVerificationRepository struct {
	db *gorm.DB
}

func NewEmailVerificationRepository(db *gorm.DB) *emailVerificationRepository {
	return &emailVerificationRepository{db: db}
}

func (repo *emailVerificationRepository) GetByEmail(email string) (result models.EmailVerification, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("emailVerificationRepository.GetByEmail failed", zap.Error(err), zap.String("email", email))
		}
	}()

	err = repo.db.Model(&models.EmailVerification{}).Where("email = ?", email).Order("generated_time desc").Find(&result).Error
	return
}

func (repo *emailVerificationRepository) Create(model models.EmailVerification) (result models.EmailVerification, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("emailVerificationRepository.Create failed", zap.Error(err), zap.Any("model", model))
		}
	}()

	err = repo.db.Model(&models.EmailVerification{}).Save(&model).Error
	if model.ID != 0 {
		result = model
	}
	return
}

func (repo *emailVerificationRepository) SetCheckedByID(id int64, isChecked bool) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("emailVerificationRepository.SetCheckedByID failed", zap.Error(err), zap.Int64("id", id), zap.Bool("isChecked", isChecked))
		}
	}()

	err = repo.db.Model(&models.EmailVerification{}).Where("id = ?", id).Update("is_checked", isChecked).Error
	return
}
