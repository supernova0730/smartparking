package cache

import "smartparking/internal/models"

type EntryHistory interface {
	Get(id int64) (result models.EntryHistory, err error)
	Set(id int64, value models.EntryHistory) (err error)
	Delete(id int64) (err error)
}
