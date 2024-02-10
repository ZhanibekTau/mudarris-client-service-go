package entityManager

import (
	"user-service-go/app/requests"
	"user-service-go/app/responses"
)

func (m *Manager) GetUstazAndClientById(payload requests.ClientAndUstazId) (*responses.ClientAndUstaz, error) {
	ustaz, err := m.services.IUstazService.GetById(payload.UstazId)
	if err != nil {
		return nil, err
	}

	client, err := m.services.IClientService.GetById(payload.ClientId)
	if err != nil {
		return nil, err
	}

	ustazStruct := responses.UstazStruct{
		Name:           ustaz.Name,
		Email:          ustaz.Email,
		Lastname:       ustaz.Lastname,
		Degree:         ustaz.Degree,
		University:     ustaz.University,
		AdditionalInfo: ustaz.AdditionalInfo,
		Experience:     ustaz.Experience,
		Phone:          ustaz.Phone,
	}

	clientStruct := responses.ClientStruct{
		Name:     client.Name,
		Lastname: client.Lastname,
		Email:    client.Email,
		Phone:    client.Phone,
	}

	resp := responses.ClientAndUstaz{
		Ustaz:  ustazStruct,
		Client: clientStruct,
	}

	return &resp, nil
}
