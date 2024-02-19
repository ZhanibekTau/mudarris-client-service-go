package services

import (
	"user-service-go/app/repositories"
	structures "user-service-go/config/configStruct"
)

type EmailService struct {
	repo repositories.EmailRepo
}

func NewEmailService(repo repositories.EmailRepo) *EmailService {
	return &EmailService{repo: repo}
}

func (e *EmailService) SendEmail(to, subject, body string, appData *structures.AppData) error {
	return e.repo.SendEmail(to, subject, body, appData)
}
