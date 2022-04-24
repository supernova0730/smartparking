package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"smartparking/internal/models"
	"smartparking/pkg/logger"
)

type parkingPlaceRepository struct {
	db *gorm.DB
}

func NewParkingPlaceRepository(db *gorm.DB) *parkingPlaceRepository {
	return &parkingPlaceRepository{db: db}
}

func (repo *parkingPlaceRepository) GetByID(id int64) (result models.ParkingPlace, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingPlaceRepository.GetByID failed", zap.Error(err), zap.Int64("id", id))
		}
	}()

	err = repo.db.Model(&models.ParkingPlace{}).Preload("ParkingZone").Where("id = ?", id).First(&result).Error
	return
}

func (repo *parkingPlaceRepository) GetAll() (result []models.ParkingPlace, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingPlaceRepository.GetAll failed", zap.Error(err))
		}
	}()

	err = repo.db.Model(&models.ParkingPlace{}).Preload("ParkingZone").Find(&result).Error
	return
}

func (repo *parkingPlaceRepository) GetAllBy(where models.ParkingPlace) (result []models.ParkingPlace, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingPlaceRepository.GetAllBy failed", zap.Error(err), zap.Any("where", where))
		}
	}()

	err = repo.db.Model(&models.ParkingPlace{}).Preload("ParkingZone").Where(&where).Find(&result).Error
	return
}

func (repo *parkingPlaceRepository) Create(model models.ParkingPlace) (result models.ParkingPlace, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingPlaceRepository.Create failed", zap.Error(err), zap.Any("model", model))
		}
	}()

	err = repo.db.Model(&models.ParkingPlace{}).Create(&model).Error
	if err != nil {
		return
	}

	return repo.GetByID(model.ID)
}

func (repo *parkingPlaceRepository) Update(id int64, model models.ParkingPlace) (result models.ParkingPlace, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingPlaceRepository.Update failed", zap.Error(err), zap.Int64("id", id), zap.Any("model", model))
		}
	}()

	err = repo.db.Model(&models.ParkingPlace{}).Where("id = ?", id).Updates(&model).Error
	if err != nil {
		return
	}

	return repo.GetByID(id)
}

func (repo *parkingPlaceRepository) SetIsBusy(id int64, isBusy bool) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingPlaceRepository.SetIsBusy failed", zap.Error(err), zap.Int64("id", id), zap.Bool("isBusy", isBusy))
		}
	}()

	err = repo.db.Model(&models.ParkingPlace{}).Where("id = ?", id).Update("is_busy", isBusy).Error
	return
}

func (repo *parkingPlaceRepository) Delete(id int64) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingPlaceRepository.Delete failed", zap.Error(err), zap.Int64("id", id))
		}
	}()

	err = repo.db.Model(&models.ParkingPlace{}).Where("id = ?", id).Delete(&models.ParkingPlace{}).Error
	return
}
