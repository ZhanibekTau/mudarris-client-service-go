package services

import (
	"client-service-go/app/repositories"
	structures "client-service-go/config/configStruct"
	"client-service-go/model"
)

type IClientService interface {
	Create(client *model.Client) (*model.Client, error)
	Login(email, password string, appData *structures.AppData) (*model.Client, error)
	Update(client model.Client) (*model.Client, error)
	GetById(id int) (*model.Client, error)
	ValidateUser(client *model.Client, appData *structures.AppData) (string, error)
	generateSessionToken(client *model.Client, appData *structures.AppData) (string, error)
}

type Service struct {
	IClientService
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		IClientService: NewClientService(repos.ClientRepo),
	}
}
