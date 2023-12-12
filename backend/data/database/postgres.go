package database

import (
	"backend/data/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var databaseClient *gorm.DB

func InitDatabase(config *config.Config) error {
	var err error
	connectionDetails := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Europe/Bucharest",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password,
		config.Database.DbName, config.Database.SSLMode)

	databaseClient, err = gorm.Open(postgres.Open(connectionDetails), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDatabase, _ := databaseClient.DB()
	err = sqlDatabase.Ping()
	if err != nil {
		return err
	}

	sqlDatabase.SetMaxIdleConns(config.Database.MaxIdleConnections)
	sqlDatabase.SetMaxOpenConns(config.Database.MaxOpenConnections)
	sqlDatabase.SetConnMaxLifetime(config.Database.ConnMaxLifetime * time.Minute)

	log.Println("Database connection established successfully.")
	return nil
}

func GetDatabase() *gorm.DB {
	return databaseClient
}

func CloseDatabase() {
	connection, _ := databaseClient.DB()
	err := connection.Close()
	if err != nil {
		log.Println("Error while closing database connection: " + err.Error())
		return
	}
}
