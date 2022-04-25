package service

import (
	"smartparking/internal/dtos"
	"smartparking/internal/models"
)

type EntryHistory interface {
	GetByID(id int64) (result models.EntryHistory, err error)
	GetAll() (result models.ListEntryHistory, err error)
	GetAllBy(where models.EntryHistory) (result models.ListEntryHistory, err error)
	GetAllByClientIDAndFilter(clientID int64, filter dtos.EntryHistoryFilter) (result models.ListEntryHistory, err error)
	Create(model models.EntryHistory) (result models.EntryHistory, err error)
	Update(id int64, model models.EntryHistory) (result models.EntryHistory, err error)
	Delete(id int64) (err error)
}
