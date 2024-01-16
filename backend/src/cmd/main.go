package main

import (
	"github.com/Stefan923/go-estate-market/api"
	"github.com/Stefan923/go-estate-market/config"
	"github.com/Stefan923/go-estate-market/data/database"
	"github.com/Stefan923/go-estate-market/data/database/migration"
	"log"
)

func main() {
	appConfig := config.GetConfig()

	err := database.InitDatabase(appConfig)
	defer database.CloseDatabase()
	if err != nil {
		log.Println("Error while initializing database: ", err.Error())
	}
	migration.Run()

	api.StartServer(appConfig)
}
