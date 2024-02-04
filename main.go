package main

import (
	"context"
	"github.com/exgamer/go-rest-sdk/pkg/modules/j/jStructures"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"time"
	_ "time/tzdata"
	"user-service-go/app/entityManager"
	"user-service-go/app/handlers"
	"user-service-go/app/repositories"
	"user-service-go/app/services"
	"user-service-go/config"
	structures "user-service-go/config/configStruct"
	gormSql "user-service-go/database/gorm"
	"user-service-go/server"
)

func init() {
	location, err := time.LoadLocation("Asia/Almaty")
	if err != nil {
		logrus.Fatalf("error loading location: %v", err.Error())
	}

	time.Local = location
}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	appConfig, dbConfig, redisConfig, restConfig, tokenConfig, rabbitConfig, err := config.InitConfig()
	if err != nil {
		logrus.Fatalf("error database connection: %v", err.Error())
	}

	appData := &structures.AppData{
		AppConfig:   appConfig,
		DbConfig:    dbConfig,
		RedisConfig: redisConfig,
		RestConfig:  restConfig,
		RequestData: &jStructures.RequestData{
			ServiceName: appConfig.Name,
		},
		TokenConfig:  tokenConfig,
		RabbitConfig: rabbitConfig,
	}

	db, err := gormSql.NewGormSqlDB(dbConfig)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer closeDbConnection(db)

	repos := repositories.NewRepository(db)
	service := services.NewService(repos)
	manager := entityManager.NewManager(service, restConfig)
	handler := handlers.NewHandler(manager, appData)

	srv := new(server.Server)

	if err := srv.Run(appConfig.ServerAddress, handler.Init()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

func closeDbConnection(db *gorm.DB) {
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
