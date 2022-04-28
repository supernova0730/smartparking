package email

import (
	"net/smtp"
)

type Manager interface {
	Send(to, message string) error
}

type manager struct {
	conf Config
}

func NewManager(conf Config) *manager {
	return &manager{conf: conf}
}

func (m *manager) Send(to, message string) error {
	address := m.conf.SMTPHost + ":" + m.conf.SMTPPort
	auth := smtp.PlainAuth("", m.conf.Sender, m.conf.Password, m.conf.SMTPHost)
	return smtp.SendMail(address, auth, m.conf.Sender, []string{to}, []byte(message))
}
