package service

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"smartparking/internal/apiError"
	"smartparking/internal/interfaces/manager"
	"smartparking/internal/models"
	"smartparking/pkg/logger"
)

type parkingZoneService struct {
	m manager.Manager
}

func NewParkingZoneService(m manager.Manager) *parkingZoneService {
	return &parkingZoneService{m: m}
}

func (s *parkingZoneService) GetByID(id int64) (result models.ParkingZone, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingZoneService.GetByID", zap.Error(err), zap.Int64("id", id))
		}
	}()

	result, err = s.m.Repository().ParkingZone().GetByID(id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = apiError.Throw(apiError.ParkingZoneNotFound)
	}
	return
}

func (s *parkingZoneService) GetAll() (result []models.ParkingZone, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingZoneService.GetAll", zap.Error(err))
		}
	}()

	return s.m.Repository().ParkingZone().GetAll()
}

func (s *parkingZoneService) Create(model models.ParkingZone) (result models.ParkingZone, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingZoneService.Create", zap.Error(err), zap.Any("model", model))
		}
	}()

	return s.m.Repository().ParkingZone().Create(model)
}

func (s *parkingZoneService) Update(id int64, model models.ParkingZone) (result models.ParkingZone, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingZoneService.Update", zap.Error(err), zap.Int64("id", id), zap.Any("model", model))
		}
	}()

	result, err = s.m.Repository().ParkingZone().Update(id, model)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = apiError.Throw(apiError.ParkingZoneNotFound)
	}
	return
}

func (s *parkingZoneService) Delete(id int64) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingZoneService.Delete", zap.Error(err), zap.Int64("id", id))
		}
	}()

	return s.m.Repository().ParkingZone().Delete(id)
}
