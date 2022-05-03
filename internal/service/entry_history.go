package service

import (
	"errors"
	"github.com/bradfitz/gomemcache/memcache"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"smartparking/internal/apiError"
	"smartparking/internal/dtos"
	"smartparking/internal/interfaces/manager"
	"smartparking/internal/models"
	"smartparking/pkg/logger"
)

type entryHistoryService struct {
	m manager.Manager
}

func NewEntryHistoryService(m manager.Manager) *entryHistoryService {
	return &entryHistoryService{m: m}
}

func (s *entryHistoryService) GetByID(id int64) (result models.EntryHistory, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("entryHistoryService.GetByID", zap.Error(err), zap.Int64("id", id))
		}
	}()

	result, err = s.m.Cache().EntryHistory().Get(id)
	if err != nil {
		if !errors.Is(err, memcache.ErrCacheMiss) {
			return
		}

		result, err = s.m.Repository().EntryHistory().GetByID(id)
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			err = apiError.Throw(apiError.EntryHistoryNotFound)
			return
		}

		err = s.m.Cache().EntryHistory().Set(id, result)
		if err != nil {
			return
		}
	}
	return
}

func (s *entryHistoryService) GetAll() (result models.ListEntryHistory, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("entryHistoryService.GetAll", zap.Error(err))
		}
	}()

	return s.m.Repository().EntryHistory().GetAll()
}

func (s *entryHistoryService) GetAllBy(where models.EntryHistory) (result models.ListEntryHistory, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("entryHistoryService.GetAllBy", zap.Error(err), zap.Any("where", where))
		}
	}()

	return s.m.Repository().EntryHistory().GetAllBy(where)
}

func (s *entryHistoryService) GetAllByClientIDAndFilter(clientID int64, filter dtos.EntryHistoryFilter) (result models.ListEntryHistory, total int64, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("entryHistoryService.GetAllByClientID", zap.Error(err), zap.Int64("clientID", clientID))
		}
	}()

	return s.m.Repository().EntryHistory().GetAllByClientIDAndFilter(clientID, filter)
}

func (s *entryHistoryService) Create(model models.EntryHistory) (result models.EntryHistory, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("entryHistoryService.Create", zap.Error(err), zap.Any("model", model))
		}
	}()

	result, err = s.m.Repository().EntryHistory().Create(model)
	if err != nil {
		return
	}

	err = s.m.Cache().EntryHistory().Set(result.ID, model)
	if err != nil {
		return
	}

	return
}

func (s *entryHistoryService) Update(id int64, model models.EntryHistory) (result models.EntryHistory, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("entryHistoryService.Update", zap.Error(err), zap.Int64("id", id), zap.Any("model", model))
		}
	}()

	result, err = s.m.Repository().EntryHistory().Update(id, model)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = apiError.Throw(apiError.EntryHistoryNotFound)
		}
		return
	}

	err = s.m.Cache().EntryHistory().Set(id, model)
	if err != nil {
		return
	}

	return
}

func (s *entryHistoryService) Delete(id int64) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("entryHistoryService.Delete", zap.Error(err), zap.Int64("id", id))
		}
	}()

	err = s.m.Repository().EntryHistory().Delete(id)
	if err != nil {
		return
	}

	err = s.m.Cache().EntryHistory().Delete(id)
	if err != nil {
		return
	}

	return
}
