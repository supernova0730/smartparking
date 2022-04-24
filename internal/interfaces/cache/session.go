package cache

import "smartparking/internal/models"

type Session interface {
	Get(id int64) (result models.Session, err error)
	Set(id int64, value models.Session) (err error)
}
