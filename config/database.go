package config

import (
	"fmt"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseConnection(configuration Config) *gorm.DB {
	dbHost := configuration.Get("DB_HOST")
	dbName := configuration.Get("DB_NAME")
	dbUser := configuration.Get("DB_USER")
	dbPassword := configuration.Get("DB_PASSWORD")
	dbPort, err := strconv.Atoi(configuration.Get("DB_PORT"))
	if err != nil {
		panic("failed to convert port")
	}

	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	database, err := gorm.Open(mysql.Open(url))
	if err != nil {
		panic("failed to connect database")
	}
	return database
}
