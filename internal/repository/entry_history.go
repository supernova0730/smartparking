package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"smartparking/internal/dtos"
	"smartparking/internal/models"
	"smartparking/pkg/logger"
)

type entryHistoryRepository struct {
	db *gorm.DB
}

func NewEntryHistoryRepository(db *gorm.DB) *entryHistoryRepository {
	return &entryHistoryRepository{db: db}
}

func (repo *entryHistoryRepository) GetByID(id int64) (result models.EntryHistory, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("entryHistoryRepository.GetByID failed", zap.Error(err), zap.Int64("id", id))
		}
	}()

	err = repo.db.
		Model(&models.EntryHistory{}).
		Preload("Car").
		Preload("Car.Client").
		Preload("ParkingZone").
		Where("id = ?", id).
		First(&result).
		Error
	return
}

func (repo *entryHistoryRepository) GetAll() (result []models.EntryHistory, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("entryHistoryRepository.GetAll failed", zap.Error(err))
		}
	}()

	err = repo.db.
		Model(&models.EntryHistory{}).
		Preload("Car").
		Preload("Car.Client").
		Preload("ParkingZone").
		Find(&result).
		Error
	return
}

func (repo *entryHistoryRepository) GetAllBy(where models.EntryHistory) (result []models.EntryHistory, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("entryHistoryRepository.GetAllBy failed", zap.Error(err), zap.Any("where", where))
		}
	}()

	err = repo.db.
		Model(&models.EntryHistory{}).
		Preload("Car").
		Preload("Car.Client").
		Preload("ParkingZone").
		Where(&where).
		Find(&result).
		Error
	return
}

func (repo *entryHistoryRepository) GetAllByClientIDAndFilter(clientID int64, filter dtos.EntryHistoryFilter) (result []models.EntryHistory, total int64, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("entryHistoryRepository.GetAllByClientID failed", zap.Error(err), zap.Int64("clientID", clientID))
		}
	}()

	err = repo.db.
		Model(&models.EntryHistory{}).
		Preload("ParkingZone").
		Joins("Car").
		Where("client_id = ?", clientID).
		Scopes(func(db *gorm.DB) *gorm.DB {
			if filter.CarID != 0 {
				db = db.Where("car_id = ?", filter.CarID)
			}
			if filter.ParkingZoneID != 0 {
				db = db.Where("parking_zone_id = ?", filter.ParkingZoneID)
			}
			if !filter.DateFrom.IsZero() && !filter.DateTo.IsZero() {
				db = db.Where("time BETWEEN ? AND ?", filter.DateFrom, filter.DateTo)
			}
			return db
		}).
		Count(&total).
		Order("time desc").
		Offset(filter.Page * filter.Size).
		Limit(filter.Size).
		Find(&result).
		Error
	return
}

func (repo *entryHistoryRepository) Create(model models.EntryHistory) (result models.EntryHistory, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("entryHistoryRepository.Create failed", zap.Error(err), zap.Any("model", model))
		}
	}()

	err = repo.db.Model(&models.EntryHistory{}).Create(&model).Error
	if err != nil {
		return
	}

	return repo.GetByID(model.ID)
}

func (repo *entryHistoryRepository) Update(id int64, model models.EntryHistory) (result models.EntryHistory, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("entryHistoryRepository.Update failed", zap.Error(err), zap.Int64("id", id), zap.Any("model", model))
		}
	}()

	err = repo.db.Model(&models.EntryHistory{}).Where("id = ?", id).Updates(&model).Error
	if err != nil {
		return
	}

	return repo.GetByID(id)
}

func (repo *entryHistoryRepository) Delete(id int64) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("entryHistoryRepository.Delete failed", zap.Error(err), zap.Int64("id", id))
		}
	}()

	err = repo.db.Model(&models.EntryHistory{}).Where("id = ?", id).Delete(&models.EntryHistory{}).Error
	return
}
