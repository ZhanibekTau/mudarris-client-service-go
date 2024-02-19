package structures

import (
	structures2 "github.com/exgamer/go-rest-sdk/pkg/config/structures"
	"github.com/exgamer/go-rest-sdk/pkg/modules/j/jStructures"
)

type AppData struct {
	RequestData    *jStructures.RequestData
	AppConfig      *structures2.AppConfig
	DbConfig       *structures2.DbConfig
	RedisConfig    *structures2.RedisConfig
	RestConfig     *RestConfig
	TokenConfig    *TokenConfig
	RabbitConfig   *RabbitConfig
	TelegramConfig *TelegramConfig
	EmailConfig    *EmailConfig
}
