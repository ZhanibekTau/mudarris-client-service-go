package entityManager

import "user-service-go/model"

func (m *Manager) Create(client *model.Client) (*model.Client, error) {
	return m.services.IClientService.Create(client)
}

func (m *Manager) Update(client model.Client) (*model.Client, error) {
	return m.services.IClientService.Update(client)
}

func (m *Manager) GetById(id int) (*model.Client, error) {
	return m.services.IClientService.GetById(id)
}
