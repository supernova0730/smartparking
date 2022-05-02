package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"smartparking/internal/models"
	"smartparking/pkg/logger"
)

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) *ticketRepository {
	return &ticketRepository{db: db}
}

func (repo *ticketRepository) GetByID(id int64) (result models.Ticket, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("ticketRepository.GetByID failed", zap.Int64("id", id))
		}
	}()

	err = repo.db.
		Model(&models.Ticket{}).
		Preload("Client").
		Preload("ParkingPlace").
		Preload("ParkingPlace.ParkingZone").
		Where("id = ?", id).
		First(&result).
		Error
	return
}

func (repo *ticketRepository) GetBy(where models.Ticket) (result models.Ticket, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("ticketRepository.GetBy failed", zap.Error(err), zap.Any("where", where))
		}
	}()

	err = repo.db.
		Model(&models.Ticket{}).
		Preload("Client").
		Preload("ParkingPlace").
		Preload("ParkingPlace.ParkingZone").
		Where(&where).
		First(&result).
		Error
	return
}

func (repo *ticketRepository) GetByParkingZoneClientIDs(parkingZoneID, clientID int64) (result models.Ticket, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("ticketRepository.GetByParkingZoneCarClientIDs failed",
				zap.Error(err),
				zap.Int64("parkingZoneID", parkingZoneID),
				zap.Int64("clientID", clientID),
			)
		}
	}()

	err = repo.db.
		Select("*").
		Table("ticket AS t").
		Joins("LEFT JOIN parking_place pp ON pp.id = t.parking_place_id").
		Where("pp.parking_zone_id = ?", parkingZoneID).
		Where("t.client_id = ?", clientID).
		Where("not pp.is_busy").
		First(&result).
		Error
	return
}

func (repo *ticketRepository) GetAll() (result []models.Ticket, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("ticketRepository.GetAll failed", zap.Error(err))
		}
	}()

	err = repo.db.
		Model(&models.Ticket{}).
		Preload("Client").
		Preload("ParkingPlace").
		Preload("ParkingPlace.ParkingZone").
		Find(&result).
		Error
	return
}

func (repo *ticketRepository) GetAllBy(where models.Ticket) (result []models.Ticket, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("ticketRepository.GetAllBy failed", zap.Error(err), zap.Any("where", where))
		}
	}()

	err = repo.db.
		Model(&models.Ticket{}).
		Preload("Client").
		Preload("ParkingPlace").
		Preload("ParkingPlace.ParkingZone").
		Where(&where).
		Find(&result).
		Error
	return
}

func (repo *ticketRepository) GetAllExpired() (result []models.Ticket, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("ticketRepository.GetAllExpired failed", zap.Error(err))
		}
	}()

	err = repo.db.
		Model(&models.Ticket{}).
		Preload("Client").
		Preload("ParkingPlace").
		Preload("ParkingPlace.ParkingZone").
		Where("expires_at < now()").
		Find(&result).
		Error
	return
}

func (repo *ticketRepository) Create(model models.Ticket) (result models.Ticket, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("ticketRepository.Create failed", zap.Error(err), zap.Any("model", model))
		}
	}()

	err = repo.db.Model(&models.Ticket{}).Create(&model).Error
	if err != nil {
		return
	}

	return repo.GetByID(model.ID)
}

func (repo *ticketRepository) Update(id int64, model models.Ticket) (result models.Ticket, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("ticketRepository.Update failed", zap.Error(err), zap.Int64("id", id), zap.Any("model", model))
		}
	}()

	err = repo.db.Model(&models.Ticket{}).Where("id = ?", id).Updates(&model).Error
	if err != nil {
		return
	}

	return repo.GetByID(id)
}

func (repo *ticketRepository) Delete(id int64) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("ticketRepository.Delete failed", zap.Error(err), zap.Int64("id", id))
		}
	}()

	err = repo.db.Model(&models.Ticket{}).Where("id = ?", id).Delete(&models.Ticket{}).Error
	return
}
