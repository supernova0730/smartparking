package manager

import "smartparking/internal/interfaces/controller"

type Controller interface {
	Auth() controller.Auth
	Car() controller.Car
	Client() controller.Client
	Core() controller.Core
	EntryHistory() controller.EntryHistory
	Image() controller.Image
	ParkingPlace() controller.ParkingPlace
	ParkingZone() controller.ParkingZone
	Personal() controller.Personal
	Tax() controller.Tax
	Ticket() controller.Ticket
}
