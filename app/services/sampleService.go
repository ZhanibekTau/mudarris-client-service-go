package services

import (
	"client-service-go/app/repositories"
)

type SampleService struct {
	repo repositories.SampleRepo
}

func NewSampleService(repo repositories.SampleRepo) *SampleService {
	return &SampleService{repo: repo}
}
