package entityManager

import (
	"user-service-go/app/services"
	structures "user-service-go/config/configStruct"
)

type Manager struct {
	services   *services.Service
	restConfig *structures.RestConfig
}

func NewManager(services *services.Service, restConfig *structures.RestConfig) *Manager {
	return &Manager{
		services:   services,
		restConfig: restConfig,
	}
}
