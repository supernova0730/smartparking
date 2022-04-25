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

type clientService struct {
	m manager.Manager
}

func NewClientService(m manager.Manager) *clientService {
	return &clientService{m: m}
}

func (s *clientService) GetByID(id int64) (result models.Client, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("clientService.GetByID", zap.Error(err), zap.Int64("id", id))
		}
	}()

	result, err = s.m.Repository().Client().GetByID(id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = apiError.Throw(apiError.ClientNotFound)
	}
	return
}

func (s *clientService) GetByPhone(phone string) (result models.Client, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("clientService.GetByPhone", zap.Error(err), zap.String("phone", phone))
		}
	}()

	result, err = s.m.Repository().Client().GetByPhone(phone)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = apiError.Throw(apiError.ClientNotFound)
	}
	return
}

func (s *clientService) GetByEmail(email string) (result models.Client, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("clientService.GetByEmail", zap.Error(err), zap.String("email", email))
		}
	}()

	result, err = s.m.Repository().Client().GetByEmail(email)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = apiError.Throw(apiError.ClientNotFound)
	}
	return
}

func (s *clientService) GetAll() (result models.ListClient, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("clientService.GetAll", zap.Error(err))
		}
	}()

	return s.m.Repository().Client().GetAll()
}

func (s *clientService) Create(model models.Client) (result models.Client, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("clientService.Create", zap.Error(err), zap.Any("model", model))
		}
	}()

	return s.m.Repository().Client().Create(model)
}

func (s *clientService) Update(id int64, model models.Client) (result models.Client, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("clientService.Update", zap.Error(err), zap.Int64("id", id), zap.Any("model", model))
		}
	}()

	result, err = s.m.Repository().Client().Update(id, model)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = apiError.Throw(apiError.ClientNotFound)
	}
	return
}

func (s *clientService) Delete(id int64) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("clientService.Delete", zap.Error(err), zap.Int64("id", id))
		}
	}()

	return s.m.Repository().Client().Delete(id)
}
