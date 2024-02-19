package entityManager

import (
	requestsForm "user-service-go/app/requests"
	structures "user-service-go/config/configStruct"
)

func (m *Manager) ForgotPassword(payload requestsForm.ForgotPassword, appData *structures.AppData) error {
	switch payload.TypeOfUser {
	case "ustaz":
		ustaz, err := m.services.IUstazService.GetByEmail(payload.Email)
		if err != nil {
			return err
		}
		err = m.services.IEmailService.SendEmail(ustaz.Email, "Пароль для "+ustaz.Name, "Ваш пароль: "+ustaz.Password, appData)
		if err != nil {
			return err
		}
		break
	case "shakirt":
		shakirt, err := m.services.IClientService.GetByEmail(payload.Email)
		if err != nil {
			return err
		}
		err = m.services.IEmailService.SendEmail(shakirt.Email, "Пароль для "+shakirt.Name, "Ваш пароль: "+shakirt.Password, appData)
		if err != nil {
			return err
		}
		break
	}

	return nil
}
