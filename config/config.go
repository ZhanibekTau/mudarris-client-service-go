package config

import (
	"github.com/davecgh/go-spew/spew"
	structures2 "github.com/exgamer/go-rest-sdk/pkg/config/structures"
	"github.com/exgamer/go-rest-sdk/pkg/helpers/config"
	"log"
	structures "user-service-go/config/configStruct"
)

func InitConfig() (*structures2.AppConfig, *structures2.DbConfig, *structures2.RedisConfig, *structures.RestConfig, *structures.TokenConfig, *structures.RabbitConfig, *structures.TelegramConfig, error) {
	appConfig, dbConfig, err := config.InitBaseConfig()
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}

	tokenConfigInterface, err := config.InitConfig(&structures.TokenConfig{})
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}

	tokenConfig, ok := tokenConfigInterface.(*structures.TokenConfig)
	if !ok {
		log.Fatalf("Cannot init rest config. Err: %s", ok)
	}

	restConfigInterface, rErr := config.InitConfig(&structures.RestConfig{})
	if rErr != nil {
		log.Fatalf("Some error occurred. Err: %s", rErr)
	}

	restConfig, ok := restConfigInterface.(*structures.RestConfig)
	if !ok {
		log.Fatalf("Cannot init rest config. Err: %s", ok)
	}

	redisConfigInterface, rErr := config.InitConfig(&structures2.RedisConfig{})
	if rErr != nil {
		log.Fatalf("Some error occurred. Err: %s", rErr)
	}

	redisConfig, ok := redisConfigInterface.(*structures2.RedisConfig)
	if !ok {
		log.Fatalf("Cannot init rest config. Err: %s", ok)
	}

	rabbitConfigInterface, rErr := config.InitConfig(&structures.RabbitConfig{})
	if rErr != nil {
		log.Fatalf("Some error occurred. Err: %s", rErr)
	}

	rabbitConfig, ok := rabbitConfigInterface.(*structures.RabbitConfig)
	if !ok {
		log.Fatalf("Cannot init rest config. Err: %s", ok)
	}

	telegramConfigInterface, rErr := config.InitConfig(&structures.TelegramConfig{})
	if rErr != nil {
		log.Fatalf("Some error occurred. Err: %s", rErr)
	}

	telegramConfig, ok := telegramConfigInterface.(*structures.TelegramConfig)
	if !ok {
		log.Fatalf("Cannot init rest config. Err: %s", ok)
	}

	if appConfig.AppEnv != "prod" {
		spew.Dump(appConfig)
		spew.Dump(dbConfig)
		spew.Dump(restConfig)
		spew.Dump(tokenConfig)
		spew.Dump(rabbitConfig)
		spew.Dump(telegramConfig)
	}

	return appConfig, dbConfig, redisConfig, restConfig, tokenConfig, rabbitConfig, telegramConfig, nil
}
