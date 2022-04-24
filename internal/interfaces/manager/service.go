package manager

import (
	"smartparking/internal/interfaces/service"
)

type Service interface {
	Auth() service.Auth
	Car() service.Car
	Client() service.Client
	EntryHistory() service.EntryHistory
	ParkingPlace() service.ParkingPlace
	ParkingZone() service.ParkingZone
	Tax() service.Tax
	Ticket() service.Ticket
	Core() service.Core
}
