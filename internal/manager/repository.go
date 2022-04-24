package manager

import (
	"smartparking/config"
	"smartparking/internal/interfaces/repository"
	repository2 "smartparking/internal/repository"
	"sync"
)

type repositoryImpl struct {
	carRepositoryInit               sync.Once
	carRepository                   repository.Car
	clientRepositoryInit            sync.Once
	clientRepository                repository.Client
	entryHistoryRepositoryInit      sync.Once
	entryHistoryRepository          repository.EntryHistory
	parkingPlaceRepositoryInit      sync.Once
	parkingPlaceRepository          repository.ParkingPlace
	parkingZoneRepositoryInit       sync.Once
	parkingZoneRepository           repository.ParkingZone
	taxRepositoryInit               sync.Once
	taxRepository                   repository.Tax
	ticketRepositoryInit            sync.Once
	ticketRepository                repository.Ticket
	emailVerificationRepositoryInit sync.Once
	emailVerificationRepository     repository.EmailVerification
	sessionRepositoryInit           sync.Once
	sessionRepository               repository.Session
	jobRepositoryInit               sync.Once
	jobRepository                   repository.Job
}

func (rm *repositoryImpl) Car() repository.Car {
	rm.carRepositoryInit.Do(func() {
		if rm.carRepository == nil {
			rm.carRepository = repository2.NewCarRepository(config.DBConn)
		}
	})
	return rm.carRepository
}

func (rm *repositoryImpl) Client() repository.Client {
	rm.clientRepositoryInit.Do(func() {
		if rm.clientRepository == nil {
			rm.clientRepository = repository2.NewClientRepository(config.DBConn)
		}
	})
	return rm.clientRepository
}

func (rm *repositoryImpl) EntryHistory() repository.EntryHistory {
	rm.entryHistoryRepositoryInit.Do(func() {
		if rm.entryHistoryRepository == nil {
			rm.entryHistoryRepository = repository2.NewEntryHistoryRepository(config.DBConn)
		}
	})
	return rm.entryHistoryRepository
}

func (rm *repositoryImpl) ParkingPlace() repository.ParkingPlace {
	rm.parkingPlaceRepositoryInit.Do(func() {
		if rm.parkingPlaceRepository == nil {
			rm.parkingPlaceRepository = repository2.NewParkingPlaceRepository(config.DBConn)
		}
	})
	return rm.parkingPlaceRepository
}

func (rm *repositoryImpl) ParkingZone() repository.ParkingZone {
	rm.parkingZoneRepositoryInit.Do(func() {
		if rm.parkingZoneRepository == nil {
			rm.parkingZoneRepository = repository2.NewParkingZoneRepository(config.DBConn)
		}
	})
	return rm.parkingZoneRepository
}

func (rm *repositoryImpl) Tax() repository.Tax {
	rm.taxRepositoryInit.Do(func() {
		if rm.taxRepository == nil {
			rm.taxRepository = repository2.NewTaxRepository(config.DBConn)
		}
	})
	return rm.taxRepository
}

func (rm *repositoryImpl) Ticket() repository.Ticket {
	rm.ticketRepositoryInit.Do(func() {
		if rm.ticketRepository == nil {
			rm.ticketRepository = repository2.NewTicketRepository(config.DBConn)
		}
	})
	return rm.ticketRepository
}

func (rm *repositoryImpl) EmailVerification() repository.EmailVerification {
	rm.emailVerificationRepositoryInit.Do(func() {
		if rm.emailVerificationRepository == nil {
			rm.emailVerificationRepository = repository2.NewEmailVerificationRepository(config.DBConn)
		}
	})
	return rm.emailVerificationRepository
}

func (rm *repositoryImpl) Session() repository.Session {
	rm.sessionRepositoryInit.Do(func() {
		if rm.sessionRepository == nil {
			rm.sessionRepository = repository2.NewSessionRepository(config.DBConn)
		}
	})
	return rm.sessionRepository
}

func (rm *repositoryImpl) Job() repository.Job {
	rm.jobRepositoryInit.Do(func() {
		if rm.jobRepository == nil {
			rm.jobRepository = repository2.NewJobRepository(config.DBConn)
		}
	})
	return rm.jobRepository
}
