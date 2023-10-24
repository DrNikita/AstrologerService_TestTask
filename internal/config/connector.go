package config

import (
	"fmt"
	"github.com/DrNikita/AstrologerService_TestTask/internal/status"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"
)

var db *gorm.DB

func init() {
	config := GetConfigurationInstance()
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DbHost, config.DbPort, config.DbUser, config.DbPass, config.DbName, config.DbSslMode,
	)
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf(status.DbConnectionError().Error())
	}
}

func GetDBInstance() *gorm.DB {
	return db
}
