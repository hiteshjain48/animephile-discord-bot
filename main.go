package main

import (
	"fmt"
	
	"github.com/hiteshjain48/animephile-discord-bot/config"
	"github.com/hiteshjain48/animephile-discord-bot/bot"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	bot.Start()

	<- make(chan struct{})
	return
}