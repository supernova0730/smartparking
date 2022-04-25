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

type parkingPlaceService struct {
	m manager.Manager
}

func NewParkingPlaceService(m manager.Manager) *parkingPlaceService {
	return &parkingPlaceService{m: m}
}

func (s *parkingPlaceService) GetByID(id int64) (result models.ParkingPlace, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingPlaceService.GetByID", zap.Error(err), zap.Int64("id", id))
		}
	}()

	result, err = s.m.Repository().ParkingPlace().GetByID(id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = apiError.Throw(apiError.ParkingPlaceNotFound)
	}
	return
}

func (s *parkingPlaceService) GetAll() (result models.ListParkingPlace, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingPlaceService.GetAll", zap.Error(err))
		}
	}()

	return s.m.Repository().ParkingPlace().GetAll()
}

func (s *parkingPlaceService) GetAllBy(where models.ParkingPlace) (result models.ListParkingPlace, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingPlaceService.GetAllBy failed", zap.Error(err), zap.Any("where", where))
		}
	}()

	return s.m.Repository().ParkingPlace().GetAllBy(where)
}

func (s *parkingPlaceService) Create(model models.ParkingPlace) (result models.ParkingPlace, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingPlaceService.Create", zap.Error(err), zap.Any("model", model))
		}
	}()

	return s.m.Repository().ParkingPlace().Create(model)
}

func (s *parkingPlaceService) Update(id int64, model models.ParkingPlace) (result models.ParkingPlace, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingPlaceService.Update", zap.Error(err), zap.Int64("id", id), zap.Any("model", model))
		}
	}()

	result, err = s.m.Repository().ParkingPlace().Update(id, model)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = apiError.Throw(apiError.ParkingPlaceNotFound)
	}
	return
}

func (s *parkingPlaceService) Delete(id int64) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("parkingPlaceService.Delete", zap.Error(err), zap.Int64("id", id))
		}
	}()

	return s.m.Repository().ParkingPlace().Delete(id)
}
