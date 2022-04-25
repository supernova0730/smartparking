package service

import "smartparking/internal/models"

type ParkingPlace interface {
	GetByID(id int64) (result models.ParkingPlace, err error)
	GetAll() (result models.ListParkingPlace, err error)
	GetAllBy(where models.ParkingPlace) (result models.ListParkingPlace, err error)
	Create(model models.ParkingPlace) (result models.ParkingPlace, err error)
	Update(id int64, model models.ParkingPlace) (result models.ParkingPlace, err error)
	Delete(id int64) (err error)
}
