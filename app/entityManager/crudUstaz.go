package entityManager

import "user-service-go/model"

func (m *Manager) CreateUstaz(ustaz *model.Ustaz) (*model.Ustaz, error) {
	return m.services.IUstazService.Create(ustaz)
}

func (m *Manager) UpdateUstaz(ustaz model.Ustaz) (*model.Ustaz, error) {
	return m.services.IUstazService.Update(ustaz)
}

func (m *Manager) GetByIdUstaz(id int) (*model.Ustaz, error) {
	return m.services.IUstazService.GetById(id)
}
