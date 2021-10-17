package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"matheusmv.com/go-bookmanager-api/config"
)

func getConnectionString(config config.Configurations) string {
	user := config.Database.User
	password := config.Database.Password
	host := config.Database.Host
	port := config.Database.Port
	database := config.Database.Name

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)

	return connectionString
}

func Connect(config config.Configurations) (*gorm.DB, error) {
	URI := getConnectionString(config)

	db, err := gorm.Open(mysql.Open(URI), &gorm.Config{})

	if err != nil {
		return db, err
	}

	return db, nil
}
