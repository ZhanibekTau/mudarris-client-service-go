package services

import (
	"user-service-go/app/repositories"
	structures "user-service-go/config/configStruct"
	"user-service-go/model"
)

type IClientService interface {
	Create(client *model.Client) (*model.Client, error)
	Login(email, password string, appData *structures.AppData) (*model.Client, error)
	Update(client model.Client) (*model.Client, error)
	GetById(id int) (*model.Client, error)
	ValidateUser(client *model.Client, appData *structures.AppData) (string, error)
	generateSessionToken(client *model.Client, appData *structures.AppData) (string, error)
	GetClientsByIds(ids []int) (*[]model.Client, error)
	GetByEmail(email string) (*model.Client, error)
}

type IUstazService interface {
	Create(ustaz *model.Ustaz) (*model.Ustaz, error)
	Login(email, password string, appData *structures.AppData) (*model.Ustaz, error)
	Update(ustaz model.Ustaz) (*model.Ustaz, error)
	GetById(id int) (*model.Ustaz, error)
	GetByEmail(email string) (*model.Ustaz, error)
}

type IEmailService interface {
	SendEmail(to, subject, body string, appData *structures.AppData) error
}

type Service struct {
	IClientService
	IUstazService
	IEmailService
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		IClientService: NewClientService(repos.ClientRepo),
		IUstazService:  NewUstazService(repos.UstazRepo),
		IEmailService:  NewEmailService(repos.EmailRepo),
	}
}
