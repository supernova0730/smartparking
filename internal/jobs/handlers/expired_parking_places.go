package handlers

import (
	"go.uber.org/zap"
	"smartparking/internal/interfaces/manager"
	"smartparking/internal/models"
	"smartparking/pkg/logger"
)

type ExpiredParkingPlaces struct {
	m   manager.Manager
	job models.Job
}

func NewExpiredParkingPlaces(m manager.Manager, job models.Job) *ExpiredParkingPlaces {
	return &ExpiredParkingPlaces{
		m:   m,
		job: job,
	}
}

func (h *ExpiredParkingPlaces) GetSchedule() string {
	return h.job.Schedule
}

func (h *ExpiredParkingPlaces) Do() {
	_ = h.m.Repository().Job().SetIsRunning(h.job.ID, true)
	defer func() {
		_ = h.m.Repository().Job().SetIsRunning(h.job.ID, false)
	}()

	expiredTickets, err := h.m.Repository().Ticket().GetAllExpired()
	if err != nil {
		logger.Log.Error("ExpiredParkingPlaces job failed", zap.Error(err))
		return
	}

	for _, expiredTicket := range expiredTickets {
		var parkingPlace models.ParkingPlace

		parkingPlace, err = h.m.Repository().ParkingPlace().GetByID(expiredTicket.ParkingPlaceID)
		if err != nil {
			logger.Log.Error("ExpiredParkingPlaces job failed", zap.Error(err))
		}

		if parkingPlace.IsBusy {
			err = h.m.Repository().ParkingPlace().SetIsBusy(parkingPlace.ID, false)
			if err != nil {
				logger.Log.Error("ExpiredParkingPlaces job failed", zap.Error(err))
			}
		}
	}
}
