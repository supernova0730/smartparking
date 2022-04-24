package repository

import (
	"smartparking/internal/dtos"
	"smartparking/internal/models"
)

type EntryHistory interface {
	GetByID(id int64) (result models.EntryHistory, err error)
	GetAll() (result []models.EntryHistory, err error)
	GetAllBy(where models.EntryHistory) (result []models.EntryHistory, err error)
	GetAllByClientIDAndFilter(clientID int64, filter dtos.EntryHistoryFilter) (result []models.EntryHistory, err error)
	Create(model models.EntryHistory) (result models.EntryHistory, err error)
	Update(id int64, model models.EntryHistory) (result models.EntryHistory, err error)
	Delete(id int64) (err error)
}
