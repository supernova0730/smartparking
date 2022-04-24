package models

import "time"

type Session struct {
	ID           int64     `gorm:"column:id"`
	ClientID     int64     `gorm:"column:client_id"`
	RefreshToken string    `gorm:"column:refresh_token"`
	ExpiresAt    time.Time `gorm:"column:expires_at"`
}

func (Session) TableName() string {
	return "session"
}

func (s Session) IsExpired() bool {
	return s.ExpiresAt.Before(time.Now())
}
