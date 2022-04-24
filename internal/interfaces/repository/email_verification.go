package repository

import "smartparking/internal/models"

type EmailVerification interface {
	GetByEmail(email string) (result models.EmailVerification, err error)
	Create(model models.EmailVerification) (result models.EmailVerification, err error)
	SetCheckedByID(id int64, isChecked bool) (err error)
}
