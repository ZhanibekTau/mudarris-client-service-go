package services

import (
	"client-service-go/app/repositories"
)

type ISampleService interface {
}

type Service struct {
	ISampleService
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		ISampleService: NewSampleService(repos.SampleRepo),
	}
}
