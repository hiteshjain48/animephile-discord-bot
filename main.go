package main

import (
	"github.com/hiteshjain48/animephile-discord-bot/bot"
	"github.com/hiteshjain48/animephile-discord-bot/config"
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

	bot.Start()

	// <- make(chan struct{})
	// return
	select {}
}
