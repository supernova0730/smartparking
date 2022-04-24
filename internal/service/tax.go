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

type taxService struct {
	m manager.Manager
}

func NewTaxService(m manager.Manager) *taxService {
	return &taxService{m: m}
}

func (s *taxService) GetByID(id int64) (result models.Tax, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("taxService.GetByID", zap.Error(err), zap.Int64("id", id))
		}
	}()

	result, err = s.m.Repository().Tax().GetByID(id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = apiError.Throw(apiError.TaxNotFound)
	}
	return
}

func (s *taxService) GetAll() (result []models.Tax, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("taxService.GetAll", zap.Error(err))
		}
	}()

	return s.m.Repository().Tax().GetAll()
}

func (s *taxService) GetAllOrderByPrice() (result []models.Tax, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("taxService.GetAllOrderByPrice", zap.Error(err))
		}
	}()

	return s.m.Repository().Tax().GetAllOrderByPrice()
}

func (s *taxService) Create(model models.Tax) (result models.Tax, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("taxService.Create", zap.Error(err), zap.Any("model", model))
		}
	}()

	return s.m.Repository().Tax().Create(model)
}

func (s *taxService) Update(id int64, model models.Tax) (result models.Tax, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("taxService.Update", zap.Error(err), zap.Int64("id", id), zap.Any("model", model))
		}
	}()

	result, err = s.m.Repository().Tax().Update(id, model)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = apiError.Throw(apiError.TaxNotFound)
	}
	return
}

func (s *taxService) Delete(id int64) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("taxService.Delete", zap.Error(err), zap.Int64("id", id))
		}
	}()

	return s.m.Repository().Tax().Delete(id)
}
