package bot

import (
	"fmt"
	"strings"

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
	message := string(msg.Content)
	messageSplit := strings.Split(message, " ")
	fmt.Println(message)
	fmt.Println(string(messageSplit[0]))
	var anime []string
	if messageSplit[1] == "subscribe" {
		for i := 2; i < len(messageSplit); i++ {
			anime = append(anime, messageSplit[i])
		}
		_, _ = session.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("subscribed to %v", anime))
	}
	fmt.Println(anime)
	fmt.Println(len(anime))
}