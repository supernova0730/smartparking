package service

import (
	"bytes"
	"go.uber.org/zap"
	"smartparking/internal/apiError"
	"smartparking/internal/interfaces/manager"
	"smartparking/internal/models"
	"smartparking/pkg/logger"
	"smartparking/pkg/tools"
	"time"
)

type coreService struct {
	m manager.Manager
}

func NewCoreService(m manager.Manager) *coreService {
	return &coreService{m: m}
}

func (s *coreService) Check(filename string, image *bytes.Buffer, parkingZoneID int64) (err error) {
	defer func() {
		if err != nil {
			logger.Log.Error("coreService.Check failed", zap.Error(err), zap.Int64("parkingZoneID", parkingZoneID))
		}
	}()

	recognitionResponse, err := s.m.Processor().PlateRecognizer().Execute(filename, image)
	if err != nil {
		return
	}

	car, err := s.m.Service().Car().GetByPlates(recognitionResponse.GetPlates())
	if err != nil {
		return
	}

	if !car.IsActive {
		err = apiError.Throw(apiError.CarIsNotActive)
		return
	}

	ticket, err := s.m.Service().Ticket().GetByParkingZoneClientIDs(parkingZoneID, car.ClientID)
	if err != nil {
		return
	}

	if ticket.IsExpired() {
		err = apiError.Throw(apiError.TicketExpired)
		return
	}

	_, err = s.m.Service().EntryHistory().Create(models.EntryHistory{
		Image:         tools.FilenameWithCurrentTime(filename),
		Time:          time.Now(),
		CarID:         car.ID,
		ParkingZoneID: parkingZoneID,
	})
	if err != nil {
		return
	}

	return nil
}
