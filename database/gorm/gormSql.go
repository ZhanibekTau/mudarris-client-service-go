package gormSql

import (
	"fmt"
	structures2 "github.com/exgamer/go-rest-sdk/pkg/config/structures"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"user-service-go/model"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func DbURL(dbConfig *structures2.DbConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Db,
	)
}

func NewGormSqlDB(dbConfig *structures2.DbConfig) (*gorm.DB, error) {
	Database, err := gorm.Open("mysql", DbURL(dbConfig))
	if err != nil {
		fmt.Println("Status:", err)
		return nil, err
	}

	Database.DB().SetMaxOpenConns(10)
	Database.DB().SetMaxIdleConns(10)

	err = Database.DB().Ping()
	if err != nil {
		return nil, err
	}
	Database.AutoMigrate(model.Client{})

	return Database, nil
}
