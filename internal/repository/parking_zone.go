package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"smartparking/internal/models"
	"smartparking/pkg/logger"
)

type parkingZoneRepository struct {
	db *gorm.DB
}

func NewParkingZoneRepository(db *gorm.DB) *parkingZoneRepository {
	return &parkingZoneRepository{db: db}
}

func (repo *parkingZoneRepository) GetByID(id int64) (result models.ParkingZone, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingZoneRepository.GetByID failed", zap.Error(err), zap.Int64("id", id))
		}
	}()

	err = repo.db.Model(&models.ParkingZone{}).Where("id = ?", id).First(&result).Error
	return
}

func (repo *parkingZoneRepository) GetAll() (result []models.ParkingZone, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingZoneRepository.GetAll failed", zap.Error(err))
		}
	}()

	err = repo.db.Model(&models.ParkingZone{}).Find(&result).Error
	return
}

func (repo *parkingZoneRepository) Create(model models.ParkingZone) (result models.ParkingZone, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingZoneRepository.Create failed", zap.Error(err), zap.Any("model", model))
		}
	}()

	err = repo.db.Model(&models.ParkingZone{}).Create(&model).Error
	if err != nil {
		return
	}

	return repo.GetByID(model.ID)
}

func (repo *parkingZoneRepository) Update(id int64, model models.ParkingZone) (result models.ParkingZone, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingZoneRepository.Update failed", zap.Error(err), zap.Int64("id", id), zap.Any("model", model))
		}
	}()

	err = repo.db.Model(&models.ParkingZone{}).Where("id = ?", id).Updates(&model).Error
	if err != nil {
		return
	}

	return repo.GetByID(id)
}

func (repo *parkingZoneRepository) Delete(id int64) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingZoneRepository.Delete failed", zap.Error(err), zap.Int64("id", id))
		}
	}()

	err = repo.db.Model(&models.ParkingZone{}).Where("id = ?", id).Delete(&models.ParkingZone{}).Error
	return
}
