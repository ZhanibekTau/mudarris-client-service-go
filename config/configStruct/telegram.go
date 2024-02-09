package structures

type TelegramConfig struct {
	TokenTelegramBot string `mapstructure:"TOKEN_TELEGRAM_BOT" json:"token_telegram_bot"`
	TelegramChatId   string `mapstructure:"TELEGRAM_CHAT_ID" json:"telegram_chat_id"`
}
