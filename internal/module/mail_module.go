package module

import (
	"auth-service-backend/config"
	"strconv"

	"gopkg.in/gomail.v2"
)

type MailModule struct {
	dialer *gomail.Dialer
	from   string
}

func NewEmailModule() *MailModule {
	port, err := strconv.Atoi(config.AppConfig.MailPort)
	if err != nil {
		panic(err)
	}

	dialer := gomail.NewDialer(
		config.AppConfig.MailHost,
		port,
		config.AppConfig.MailUser,
		config.AppConfig.MailPassword,
	)

	return &MailModule{
		dialer: dialer,
		from:   config.AppConfig.MailUser,
	}
}

func (m *MailModule) SendEmail(to, subject, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", m.from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	if err := m.dialer.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
