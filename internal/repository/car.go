package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"smartparking/internal/models"
	"smartparking/pkg/logger"
	"smartparking/pkg/tools"
	"strings"
)

type carRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) *carRepository {
	return &carRepository{db: db}
}

func (repo *carRepository) GetByID(id int64) (result models.Car, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("carRepository.GetByID failed", zap.Error(err), zap.Int64("id", id))
		}
	}()

	err = repo.db.Model(&models.Car{}).Preload("Client").Where("id = ?", id).First(&result).Error
	return
}

func (repo *carRepository) GetByPlate(plate string) (result models.Car, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("carRepository.GetByPlate failed", zap.Error(err), zap.String("plate", plate))
		}
	}()

	err = repo.db.Model(&models.Car{}).Preload("Client").Where("UPPER(plate) = UPPER(?)", plate).First(&result).Error
	return
}

func (repo *carRepository) GetByPlates(plates []string) (result models.Car, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("carRepository.GetByPlateList failed", zap.Error(err), zap.Strings("plates", plates))
		}
	}()

	plates = tools.SliceToUpper(plates)
	err = repo.db.Model(&models.Car{}).Preload("Client").Where("UPPER(plate) IN ?", plates).First(&result).Error
	return
}

func (repo *carRepository) GetAll() (result []models.Car, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("carRepository.GetAll failed", zap.Error(err))
		}
	}()

	err = repo.db.Model(&models.Car{}).Preload("Client").Find(&result).Error
	return
}

func (repo *carRepository) GetAllBy(where models.Car) (result []models.Car, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("carRepository.GetAllBy failed", zap.Error(err), zap.Any("where", where))
		}
	}()

	err = repo.db.Model(&models.Car{}).Preload("Client").Where(&where).Find(&result).Error
	return
}

func (repo *carRepository) Create(model models.Car) (result models.Car, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("carRepository.Create failed", zap.Error(err), zap.Any("model", model))
		}
	}()

	model.Plate = strings.ToUpper(model.Plate)
	err = repo.db.Model(&models.Car{}).Create(&model).Error
	if err != nil {
		return
	}

	return repo.GetByID(model.ID)
}

func (repo *carRepository) Update(id int64, model models.Car) (result models.Car, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("carRepository.Update failed", zap.Error(err), zap.Int64("id", id), zap.Any("model", model))
		}
	}()

	model.Plate = strings.ToUpper(model.Plate)
	err = repo.db.Model(&models.Car{}).Where("id = ?", id).Updates(&model).Error
	if err != nil {
		return
	}

	return repo.GetByID(id)
}

func (repo *carRepository) SetStatusByClientAndCarID(clientID, carID int64, isActive bool) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error(
				"carRepository.SetStatusByClientAndCarID failed",
				zap.Error(err),
				zap.Int64("clientID", clientID),
				zap.Any("carID", carID),
				zap.Bool("isActive", isActive),
			)
		}
	}()

	err = repo.db.
		Model(&models.Car{}).
		Where("client_id = ? and id = ?", clientID, carID).
		Update("is_active", isActive).
		Error
	return
}

func (repo *carRepository) Delete(id int64) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("carRepository.Delete failed", zap.Error(err), zap.Int64("id", id))
		}
	}()

	err = repo.db.Model(&models.Car{}).Where("id = ?", id).Delete(&models.Car{}).Error
	return
}
