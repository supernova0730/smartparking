package repository

import "smartparking/internal/models"

type Session interface {
	GetByID(id int64) (result models.Session, err error)
	GetByRefreshToken(refreshToken string) (result models.Session, err error)
	Create(model models.Session) (result models.Session, err error)
	DeleteByID(id int64) (err error)
}
