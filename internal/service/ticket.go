package service

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"smartparking/internal/apiError"
	"smartparking/internal/dtos"
	"smartparking/internal/interfaces/manager"
	"smartparking/internal/models"
	"smartparking/pkg/logger"
	"time"
)

type ticketService struct {
	m manager.Manager
}

func NewTicketService(m manager.Manager) *ticketService {
	return &ticketService{m: m}
}

func (s *ticketService) GetByID(id int64) (result models.Ticket, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("ticketService.GetByID", zap.Error(err), zap.Int64("id", id))
		}
	}()

	result, err = s.m.Repository().Ticket().GetByID(id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = apiError.Throw(apiError.TicketNotFound)
	}
	return
}

func (s *ticketService) GetBy(where models.Ticket) (result models.Ticket, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("ticketService.GetBy failed", zap.Error(err), zap.Any("where", where))
		}
	}()

	result, err = s.m.Repository().Ticket().GetBy(where)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = apiError.Throw(apiError.TicketNotFound)
	}
	return
}

func (s *ticketService) GetByParkingZoneClientIDs(parkingZoneID, clientID int64) (result models.Ticket, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error(
				"ticketService.GetByParkingZoneCarClientIDs failed",
				zap.Error(err),
				zap.Int64("parkingZoneID", parkingZoneID),
				zap.Int64("clientID", clientID),
			)
		}
	}()

	result, err = s.m.Repository().Ticket().GetByParkingZoneClientIDs(parkingZoneID, clientID)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = apiError.Throw(apiError.TicketNotFound)
	}
	return
}

func (s *ticketService) GetAll() (result []models.Ticket, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("ticketService.GetAll", zap.Error(err))
		}
	}()

	return s.m.Repository().Ticket().GetAll()
}

func (s *ticketService) GetAllBy(where models.Ticket) (result []models.Ticket, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("ticketService.GetAllBy failed", zap.Error(err), zap.Any("where", where))
		}
	}()
	return s.m.Repository().Ticket().GetAllBy(where)
}

func (s *ticketService) Create(model models.Ticket) (result models.Ticket, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("ticketService.Create", zap.Error(err), zap.Any("model", model))
		}
	}()

	return s.m.Repository().Ticket().Create(model)
}

func (s *ticketService) BuyTicket(clientID int64, buyTicket dtos.BuyTicketDTO) (result models.Ticket, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("ticketService.BuyTicket failed", zap.Error(err), zap.Any("buyTicket", buyTicket))
		}
	}()

	tax, err := s.m.Service().Tax().GetByID(buyTicket.TaxID)
	if err != nil {
		return
	}

	parkingPlace, err := s.m.Service().ParkingPlace().GetByID(buyTicket.ParkingPlaceID)
	if err != nil {
		return
	}

	model := models.Ticket{
		ExpiresAt:      time.Now().Add(tax.Duration),
		ClientID:       clientID,
		ParkingPlaceID: parkingPlace.ID,
	}

	return s.m.Service().Ticket().Create(model)
}

func (s *ticketService) Update(id int64, model models.Ticket) (result models.Ticket, err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("ticketService.Update", zap.Error(err), zap.Int64("id", id), zap.Any("model", model))
		}
	}()

	result, err = s.m.Repository().Ticket().Update(id, model)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = apiError.Throw(apiError.TicketNotFound)
	}
	return
}

func (s *ticketService) Delete(id int64) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("ticketService.Delete", zap.Error(err), zap.Int64("id", id))
		}
	}()

	return s.m.Repository().Ticket().Delete(id)
}
