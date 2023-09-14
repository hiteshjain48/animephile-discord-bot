package bot

import (
	"fmt"

	"github.com/hiteshjain48/animephile-discord-bot/config"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot "+ config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := goBot.User("@me")
	fmt.Println(user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotID = user.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
}


func messageHandler(session *discordgo.Session, msg *discordgo.MessageCreate) {
	
	if msg.Author.ID == BotID {
		return
	}
	fmt.Println(msg.Author.ID)
	if msg.Content == "subscribe" {
		_, _ = session.ChannelMessageSend(msg.ChannelID, "done!")
	}
}