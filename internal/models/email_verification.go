package models

import (
	"smartparking/internal/apiError"
	"smartparking/internal/consts"
	"strings"
	"time"
)

type EmailVerification struct {
	ID            int64     `gorm:"column:id"`
	Email         string    `gorm:"column:email"`
	Password      string    `gorm:"column:password"`
	Code          string    `gorm:"column:code"`
	IsChecked     bool      `gorm:"column:is_checked"`
	GeneratedTime time.Time `gorm:"column:generated_time"`
}

func (EmailVerification) TableName() string {
	return "email_verification"
}

func (ev EmailVerification) IsExpired() bool {
	diff := time.Now().Sub(ev.GeneratedTime)
	validDuration := 2 * time.Minute
	return diff > validDuration
}

func (ev EmailVerification) IsValid(code string) error {
	if ev.IsChecked || ev.IsExpired() {
		return apiError.Throw(apiError.EmailVerificationCodeExpired)
	}

	if strings.TrimSpace(ev.Code) == consts.DefaultOTP {
		return nil
	}

	if strings.TrimSpace(ev.Code) != strings.TrimSpace(code) {
		return apiError.Throw(apiError.EmailVerificationCodeInvalid)
	}

	return nil
}
