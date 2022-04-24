package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"smartparking/internal/models"
	"smartparking/pkg/logger"
)

type taxRepository struct {
	db *gorm.DB
}

func NewTaxRepository(db *gorm.DB) *taxRepository {
	return &taxRepository{db: db}
}

func (repo *taxRepository) GetByID(id int64) (result models.Tax, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("taxRepository.GetByID failed", zap.Error(err), zap.Int64("id", id))
		}
	}()

	err = repo.db.Model(&models.Tax{}).Where("id = ?", id).First(&result).Error
	return
}

func (repo *taxRepository) GetAll() (result []models.Tax, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("taxRepository.GetAll failed", zap.Error(err))
		}
	}()

	err = repo.db.Model(&models.Tax{}).Find(&result).Error
	return
}

func (repo *taxRepository) GetAllOrderByPrice() (result []models.Tax, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("taxRepository.GetAllOrderByPrice failed", zap.Error(err))
		}
	}()

	err = repo.db.Model(&models.Tax{}).Order("price").Find(&result).Error
	return
}

func (repo *taxRepository) Create(model models.Tax) (result models.Tax, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("taxRepository.Create failed", zap.Error(err), zap.Any("model", model))
		}
	}()

	err = repo.db.Model(&models.Tax{}).Create(&model).Error
	if err != nil {
		return
	}

	return repo.GetByID(model.ID)
}

func (repo *taxRepository) Update(id int64, model models.Tax) (result models.Tax, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("taxRepository.Update failed", zap.Error(err), zap.Int64("id", id), zap.Any("model", model))
		}
	}()

	err = repo.db.Model(&models.Tax{}).Where("id = ?", id).Updates(&model).Error
	if err != nil {
		return
	}

	return repo.GetByID(id)
}

func (repo *taxRepository) Delete(id int64) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("taxRepository.Delete failed", zap.Error(err), zap.Int64("id", id))
		}
	}()

	err = repo.db.Model(&models.Tax{}).Where("id = ?", id).Delete(&models.Tax{}).Error
	return
}
