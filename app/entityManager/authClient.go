package entityManager

import (
	"client-service-go/app/requests"
	structures "client-service-go/config/configStruct"
	"client-service-go/model"
)

func (m *Manager) Login(login requests.Login, appData *structures.AppData) (*model.Client, error) {
	ustaz, err := m.services.IClientService.Login(login.Email, login.Password, appData)
	if err != nil {
		return nil, err
	}

	return ustaz, nil
}
