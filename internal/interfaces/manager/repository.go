package manager

import "smartparking/internal/interfaces/repository"

type Repository interface {
	Car() repository.Car
	Client() repository.Client
	EntryHistory() repository.EntryHistory
	ParkingPlace() repository.ParkingPlace
	ParkingZone() repository.ParkingZone
	Tax() repository.Tax
	Ticket() repository.Ticket
	EmailVerification() repository.EmailVerification
	Session() repository.Session
	Job() repository.Job
}
