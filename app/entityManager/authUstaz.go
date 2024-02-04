package entityManager

import (
	requestsForm "user-service-go/app/requests"
	structures "user-service-go/config/configStruct"
	"user-service-go/model"
)

func (m *Manager) LoginUstaz(login requestsForm.Login, appData *structures.AppData) (*model.Ustaz, error) {
	ustaz, err := m.services.IUstazService.Login(login.Email, login.Password, appData)
	if err != nil {
		return nil, err
	}

	return ustaz, nil
}
