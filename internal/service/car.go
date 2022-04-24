package service

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"smartparking/internal/apiError"
	"smartparking/internal/consts"
	"smartparking/internal/interfaces/manager"
	"smartparking/internal/models"
	"smartparking/pkg/logger"
)

type carService struct {
	m manager.Manager
}

func NewCarService(m manager.Manager) *carService {
	return &carService{m: m}
}

func (s *carService) GetByID(id int64) (result models.Car, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("carService.GetByID", zap.Error(err), zap.Int64("id", id))
		}
	}()

	result, err = s.m.Repository().Car().GetByID(id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = apiError.Throw(apiError.CarNotFound)
	}
	return
}

func (s *carService) GetByPlate(plate string) (result models.Car, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("carService.GetByPlate", zap.Error(err), zap.String("plate", plate))
		}
	}()

	result, err = s.m.Repository().Car().GetByPlate(plate)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = apiError.Throw(apiError.CarNotFound)
	}
	return
}

func (s *carService) GetByPlates(plates []string) (result models.Car, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("carService.GetByPlates", zap.Error(err), zap.Strings("plates", plates))
		}
	}()

	return s.m.Repository().Car().GetByPlates(plates)
}

func (s *carService) GetAll() (result []models.Car, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("carService.GetAll", zap.Error(err))
		}
	}()

	return s.m.Repository().Car().GetAll()
}

func (s *carService) GetAllByClientID(clientID int64) (result []models.Car, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("carService.GetAllByClientID", zap.Error(err), zap.Int64("clientID", clientID))
		}
	}()

	return s.m.Repository().Car().GetAllBy(models.Car{
		ClientID: clientID,
	})
}

func (s *carService) Create(model models.Car) (result models.Car, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("carService.Create", zap.Error(err), zap.Any("model", model))
		}
	}()

	return s.m.Repository().Car().Create(model)
}

func (s *carService) Update(id int64, model models.Car) (result models.Car, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("carService.Update", zap.Error(err), zap.Int64("id", id), zap.Any("model", model))
		}
	}()

	result, err = s.m.Repository().Car().Update(id, model)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = apiError.Throw(apiError.CarNotFound)
	}
	return
}

func (s *carService) SetStatusByClientAndCarID(clientID, carID int64, isActive bool) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error(
				"carService.SetStatusByID",
				zap.Error(err),
				zap.Int64("clientID", clientID),
				zap.Int64("carID", carID),
				zap.Bool("isActive", isActive),
			)
		}
	}()

	car, err := s.m.Repository().Car().GetByID(carID)
	if err != nil {
		return
	}

	if car.ClientID != clientID {
		err = apiError.Throw(apiError.NotYourCar)
		return
	}

	if isActive {
		var activeCars []models.Car
		activeCars, err = s.m.Repository().Car().GetAllBy(models.Car{
			ClientID: clientID,
			IsActive: true,
		})
		if err != nil {
			return
		}

		if len(activeCars) >= consts.NUM_OF_ACTIVE_CARS {
			last := consts.NUM_OF_ACTIVE_CARS - 1
			err = s.m.Repository().Car().SetStatusByClientAndCarID(clientID, activeCars[last].ID, false)
			if err != nil {
				return
			}
		}
	}

	return s.m.Repository().Car().SetStatusByClientAndCarID(clientID, carID, isActive)
}

func (s *carService) Delete(id int64) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("carService.", zap.Error(err), zap.Int64("id", id))
		}
	}()

	return s.m.Repository().Car().Delete(id)
}
