package service

import "smartparking/internal/models"

type ParkingZone interface {
	GetByID(id int64) (result models.ParkingZone, err error)
	GetAll() (result models.ListParkingZone, err error)
	Create(model models.ParkingZone) (result models.ParkingZone, err error)
	Update(id int64, model models.ParkingZone) (result models.ParkingZone, err error)
	Delete(id int64) (err error)
}
