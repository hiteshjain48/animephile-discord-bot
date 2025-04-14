package main

import (
	"github.com/hiteshjain48/animephile-discord-bot/bot"
	"github.com/hiteshjain48/animephile-discord-bot/config"
	"github.com/hiteshjain48/animephile-discord-bot/database"
	"github.com/hiteshjain48/animephile-discord-bot/logger"
)

func main() {
	logger.Init()
	err := config.ReadConfig()
	if err != nil {
		// fmt.Println(err)
		logger.Log.Error(err)
		return
	}
	dbConfig := database.DBConfig{
		Host:     config.DBHost,
		Port:     int(config.DBPort),
		User:     config.DBUser,
		Password: config.DBPass,
		DBName:   config.DBName,
		SSLMode:  config.DBSSLMode,
	}
	db, err := database.Connect(dbConfig)
	if err != nil {
		logger.Log.Error(err)
		return
	}
	err = database.RunMigrations(db, "./database/migrations")
	if err != nil {
		return
	}
	bot.Start()

	// <- make(chan struct{})
	// return
	select {}
}
