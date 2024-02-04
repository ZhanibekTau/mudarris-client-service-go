package environment

import (
	"fmt"
	"os"
	"strconv"
	gormSql "user-service-go/database/gorm"
)

func GetDbConfigWithEnvs() (*gormSql.DBConfig, error) {
	dbHost, err := getEnv("DB_HOST")
	if err != nil {
		return nil, err
	}
	dbPort, err := getEnvInt("DB_PORT")
	if err != nil {
		return nil, err
	}
	dbName, err := getEnv("DB_NAME")
	if err != nil {
		return nil, err
	}
	dbUser, err := getEnv("DB_USER")
	if err != nil {
		return nil, err
	}
	dbPwd, err := getEnv("DB_PWD")
	if err != nil {
		return nil, err
	}

	return &gormSql.DBConfig{
		Host:     dbHost,
		Port:     dbPort,
		DBName:   dbName,
		User:     dbUser,
		Password: dbPwd,
	}, nil
}

func getEnvInt(key string) (int, error) {
	value, exists := os.LookupEnv(key)
	if exists {
		value, err := strconv.Atoi(value)
		if err != nil {
			return 0, err
		}
		return value, nil
	}

	return 0, fmt.Errorf("%s env missing", key)
}

func getEnv(key string) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}
	return "", fmt.Errorf("%s env missing", key)
}
