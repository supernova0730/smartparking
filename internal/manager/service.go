package manager

import (
	"smartparking/config"
	"smartparking/internal/interfaces/service"
	service2 "smartparking/internal/service"
	"smartparking/pkg/email"
	"smartparking/pkg/hash"
	"smartparking/pkg/jwt"
	"sync"
)

type serviceImpl struct {
	authServiceInit         sync.Once
	authService             service.Auth
	carServiceInit          sync.Once
	carService              service.Car
	clientServiceInit       sync.Once
	clientService           service.Client
	entryHistoryServiceInit sync.Once
	entryHistoryService     service.EntryHistory
	parkingPlaceServiceInit sync.Once
	parkingPlaceService     service.ParkingPlace
	parkingZoneServiceInit  sync.Once
	parkingZoneService      service.ParkingZone
	taxServiceInit          sync.Once
	taxService              service.Tax
	ticketServiceInit       sync.Once
	ticketService           service.Ticket
	coreServiceInit         sync.Once
	coreService             service.Core
}

func (sm *serviceImpl) Auth() service.Auth {
	sm.authServiceInit.Do(func() {
		if sm.authService == nil {
			sm.authService = service2.NewAuthService(
				MManager,
				jwt.NewManager(config.GlobalConfig.JWT.SecretKey),
				hash.NewManager(),
				email.NewManager(email.Config{
					Sender:   config.GlobalConfig.Email.Sender,
					Password: config.GlobalConfig.Email.Password,
					SMTPHost: config.GlobalConfig.Email.SMTPHost,
					SMTPPort: config.GlobalConfig.Email.SMTPPort,
				}),
				config.GlobalConfig.JWT.AccessTokenTTL,
				config.GlobalConfig.JWT.RefreshTokenTTL,
			)
		}
	})
	return sm.authService
}

func (sm *serviceImpl) Car() service.Car {
	sm.carServiceInit.Do(func() {
		if sm.carService == nil {
			sm.carService = service2.NewCarService(MManager)
		}
	})
	return sm.carService
}

func (sm *serviceImpl) Client() service.Client {
	sm.clientServiceInit.Do(func() {
		if sm.clientService == nil {
			sm.clientService = service2.NewClientService(MManager)
		}
	})
	return sm.clientService
}

func (sm *serviceImpl) EntryHistory() service.EntryHistory {
	sm.entryHistoryServiceInit.Do(func() {
		if sm.entryHistoryService == nil {
			sm.entryHistoryService = service2.NewEntryHistoryService(MManager)
		}
	})
	return sm.entryHistoryService
}

func (sm *serviceImpl) ParkingPlace() service.ParkingPlace {
	sm.parkingPlaceServiceInit.Do(func() {
		if sm.parkingPlaceService == nil {
			sm.parkingPlaceService = service2.NewParkingPlaceService(MManager)
		}
	})
	return sm.parkingPlaceService
}

func (sm *serviceImpl) ParkingZone() service.ParkingZone {
	sm.parkingZoneServiceInit.Do(func() {
		if sm.parkingZoneService == nil {
			sm.parkingZoneService = service2.NewParkingZoneService(MManager)
		}
	})
	return sm.parkingZoneService
}

func (sm *serviceImpl) Tax() service.Tax {
	sm.taxServiceInit.Do(func() {
		if sm.taxService == nil {
			sm.taxService = service2.NewTaxService(MManager)
		}
	})
	return sm.taxService
}

func (sm *serviceImpl) Ticket() service.Ticket {
	sm.ticketServiceInit.Do(func() {
		if sm.ticketService == nil {
			sm.ticketService = service2.NewTicketService(MManager)
		}
	})
	return sm.ticketService
}

func (sm *serviceImpl) Core() service.Core {
	sm.coreServiceInit.Do(func() {
		if sm.coreService == nil {
			sm.coreService = service2.NewCoreService(MManager)
		}
	})
	return sm.coreService
}
