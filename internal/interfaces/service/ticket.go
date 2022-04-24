package service

import (
	"smartparking/internal/dtos"
	"smartparking/internal/models"
)

type Ticket interface {
	GetByID(id int64) (result models.Ticket, err error)
	GetBy(where models.Ticket) (result models.Ticket, err error)
	GetByParkingZoneClientIDs(parkingZoneID, clientID int64) (result models.Ticket, err error)
	GetAll() (result []models.Ticket, err error)
	GetAllBy(where models.Ticket) (result []models.Ticket, err error)
	Create(model models.Ticket) (result models.Ticket, err error)
	BuyTicket(clientID int64, buyTicket dtos.BuyTicketDTO) (result models.Ticket, err error)
	Update(id int64, model models.Ticket) (result models.Ticket, err error)
	Delete(id int64) (err error)
}
