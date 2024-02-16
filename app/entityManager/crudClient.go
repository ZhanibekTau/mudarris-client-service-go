package entityManager

import (
	"user-service-go/app/responses"
	"user-service-go/model"
)

func (m *Manager) Create(client *model.Client) (*model.Client, error) {
	return m.services.IClientService.Create(client)
}

func (m *Manager) Update(client model.Client) (*model.Client, error) {
	return m.services.IClientService.Update(client)
}

func (m *Manager) GetById(id int) (*model.Client, error) {
	return m.services.IClientService.GetById(id)
}

func (m *Manager) GetClientsByIds(ids []int) (*[]responses.ClientStruct, error) {
	clients, err := m.services.IClientService.GetClientsByIds(ids)
	if err != nil {
		return nil, err
	}

	var clientsResponse []responses.ClientStruct
	for _, item := range *clients {
		result := responses.ClientStruct{
			Id:       item.Id,
			Name:     item.Name,
			Lastname: item.Lastname,
			Phone:    item.Phone,
			Email:    item.Email,
		}
		clientsResponse = append(clientsResponse, result)
	}

	return &clientsResponse, nil
}
