package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/hiteshjain48/animephile-discord-bot/config"
	"github.com/hiteshjain48/animephile-discord-bot/logger"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	logger.Init()
	var err error
	goBot, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		// fmt.Println(err.Error())
		logger.Log.Error(err.Error())
		return
	}

	user, err := goBot.User("@me")
	// fmt.Println(user)
	logger.Log.Info(fmt.Sprintf("User: %s", user))
	if err != nil {
		// fmt.Println(err.Error())
		logger.Log.Error(err.Error())
		return
	}

	BotID = user.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		// fmt.Println(err.Error())
		logger.Log.Error(err.Error())
		return
	}

	// fmt.Println("Bot is running!")
	logger.Log.Info("Bot is running!")
}

func messageHandler(session *discordgo.Session, msg *discordgo.MessageCreate) {

	if msg.Author.ID == BotID {
		return
	}
	// fmt.Println(msg.Author.ID)
	// fmt.Println(msg.Content)
	logger.Log.Info(fmt.Sprintf("Author: %s", msg.Author))
	logger.Log.Info(fmt.Sprintf("Message received: ", msg.Content))
	if !strings.HasPrefix(msg.Content, config.BotPrefix) {
		return
	}

	// message := string(msg.Content)

	//this'll work if doing from channel
	// msgSplit := strings.Split(msg.Content, " ")[1:]
	// fmt.Println(msgSplit)
	// fmt.Println(config.BotPrefix)
	// message := strings.TrimPrefix(strings.Join(msgSplit, " "), config.BotPrefix)
	// fmt.Println(message)
	// args := strings.Fields(message)

	//if direcct msg to bot then this
	message := strings.TrimPrefix(msg.Content, config.BotPrefix)
	args := strings.Fields(message)
	if len(args) == 0 {
		return
	}

	command := args[0]

	switch command {
	case "subscribe":
		if len(args) < 2 {
			session.ChannelMessageSend(msg.ChannelID, "Please specify anime to subscribe to.")
			return
		}
		animeList := args[1:]
		session.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("Subscribed to %s", strings.Join(animeList, ", ")))
	case "list":
		session.ChannelMessageSend(msg.ChannelID, "You are not subscribed to any anime yet.")
	case "help":
		helpMessage := "Available commands:\n" +
			"!subscribe [anime name] - Subscribe to anime updates\n" +
			"!list - Show your subscriptions\n" +
			"!help - Show this message"
		session.ChannelMessageSend(msg.ChannelID, helpMessage)
	default:
		session.ChannelMessageSend(msg.ChannelID, "Unknown Command")
	}
	// messageSplit := strings.Split(message, " ")
	// fmt.Println(message)
	// fmt.Println(string(messageSplit[0]))
	// var anime []string
	// if messageSplit[1] == "subscribe" {
	// 	for i := 2; i < len(messageSplit); i++ {
	// 		anime = append(anime, messageSplit[i])
	// 	}
	// 	_, _ = session.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("subscribed to %v", anime))
	// }
	// fmt.Println(anime)
	// fmt.Println(len(anime))
}
