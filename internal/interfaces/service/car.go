package service

import "smartparking/internal/models"

type Car interface {
	GetByID(id int64) (result models.Car, err error)
	GetByPlate(plate string) (result models.Car, err error)
	GetByPlates(plates []string) (result models.Car, err error)
	GetAll() (result models.ListCar, err error)
	GetAllByClientID(clientID int64) (result models.ListCar, err error)
	Create(model models.Car) (result models.Car, err error)
	Update(id int64, model models.Car) (result models.Car, err error)
	SetStatusByClientAndCarID(clientID, carID int64, isActive bool) (err error)
	Delete(id int64) (err error)
}
