package entityManager

import (
	"client-service-go/app/services"
	structures "client-service-go/config/configStruct"
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
