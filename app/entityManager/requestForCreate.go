package entityManager

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	structures "user-service-go/config/configStruct"
	"user-service-go/model"
)

func (m *Manager) RequestForCreate(ustaz *model.Ustaz, appData *structures.AppData) error {
	bot, err := tgbotapi.NewBotAPI(appData.TelegramConfig.TokenTelegramBot)
	if err != nil {
		log.Panic(err)
	}
	var text string

	text = fmt.Sprintf(
		"Запрос на регистрацию!\n\n\nИмя: %s\nФамилия: %s\nТелефон: %s\nПочта: %s\nУровень образования: %s\nУниверситет: %s\nОпыт(в годах): %.f\nДоп. информация: %s",
		ustaz.Name, ustaz.Lastname, ustaz.Phone, ustaz.Email, ustaz.Degree, ustaz.University, ustaz.Experience, ustaz.AdditionalInfo)

	chatId, err := strconv.ParseInt(appData.TelegramConfig.TelegramChatId, 10, 64)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}

	msg := tgbotapi.NewMessage(chatId, text)

	_, err = bot.Send(msg)
	if err != nil {
		log.Panic(err)
		return err
	}
	return nil

}
