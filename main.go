package main

import (
	"github.com/hiteshjain48/animephile-discord-bot/bot"
	"github.com/hiteshjain48/animephile-discord-bot/config"
	"github.com/hiteshjain48/animephile-discord-bot/database"
	"github.com/hiteshjain48/animephile-discord-bot/logger"
	"github.com/hiteshjain48/animephile-discord-bot/database/repositories"
	"github.com/hiteshjain48/animephile-discord-bot/scheduler"

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
		DBUrl: 	  config.DBUrl,	
	}
	logger.Log.Info(dbConfig.DBUrl)
	db, err := database.Connect(dbConfig)
	if err != nil {
		logger.Log.Error(err)
		return
	}
	err = database.RunMigrations(db, "./database/migrations")
	if err != nil {
		logger.Log.Error(err)
		return
	}

	userRepo := repositories.NewUserRepository(db)
    animeRepo := repositories.NewAnimeRepository(db)
	subscriptionRepo := repositories.NewSubscriptionRepository(db)
	goBot, err := bot.Start(userRepo, animeRepo, subscriptionRepo)
	if err != nil {
		logger.Log.Error(err)
		return
	}
	scheduler := scheduler.NewScheduler(animeRepo, subscriptionRepo, goBot)
	scheduler.Start()
	defer scheduler.Stop()

	// <- make(chan struct{})
	// return
	select {}
}
