package repositories

import (
	"gopkg.in/gomail.v2"
	"strconv"
	structures "user-service-go/config/configStruct"
)

type EmailRepository struct {
}

func NewEmailRepository() *EmailRepository {
	return &EmailRepository{}
}

func (e *EmailRepository) SendEmail(to, subject, body string, appData *structures.AppData) error {
	port, err := strconv.Atoi(appData.EmailConfig.EmailPort)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", appData.EmailConfig.EmailUsername)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(appData.EmailConfig.EmailHost, port, appData.EmailConfig.EmailUsername, appData.EmailConfig.EmailPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
