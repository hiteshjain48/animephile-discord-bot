package bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/hiteshjain48/animephile-discord-bot/anime"
	"github.com/hiteshjain48/animephile-discord-bot/config"
	"github.com/hiteshjain48/animephile-discord-bot/database/models"
	"github.com/hiteshjain48/animephile-discord-bot/database/repositories"
	"github.com/hiteshjain48/animephile-discord-bot/logger"
)

var BotID string
var goBot *discordgo.Session
var uRepo *repositories.UserRepository
var aRepo *repositories.AnimeRepository
var sRepo *repositories.SubscriptionRepository

func Start(userRepo *repositories.UserRepository, animeRepo *repositories.AnimeRepository, subscriptionRepo *repositories.SubscriptionRepository) {
	logger.Init()
	var err error
	uRepo = userRepo
	aRepo = animeRepo
	sRepo = subscriptionRepo
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
	schedules, err := anime.GetSchedule()
	if err != nil {
		logger.Log.Error(err)
	}
	for _, s := range schedules {
		fmt.Printf("📺 %s - Ep %d at %s\n",
			s.Media.Title.Romaji,
			s.Episode,
			time.Unix(s.AiringAt, 0).Format("15:04 MST"))
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
	logger.Log.Info(fmt.Sprintf("Message received: %s", msg.Content))
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
		animePresent, err := aRepo.List()
		animePresentLookup := make(map[string]struct{})
		for _, anime := range animePresent {
			if _, exists := animePresentLookup[anime.Title]; !exists {
				animePresentLookup[anime.Title] = struct{}{}
			}
		}
		if err != nil {
			session.ChannelMessageSend(msg.ChannelID, "Can't subscribe right now.")
			logger.Log.Error(err.Error())
			break
		}
		var user models.User
		_, err = uRepo.GetByID(msg.Author.ID)
		if err != nil {
			if err.Error() == "user not found" {
				user = models.User{
					DiscordID: msg.Author.ID,
					UserName:  msg.Author.Username,
					JoinedAt:  time.Now(),
				}
				err = uRepo.Create(user)
				if err != nil {
					session.ChannelMessageSend(msg.ChannelID, "Error creating user try again")
					logger.Log.Error(err.Error())
					break
				}
			} else {
				session.ChannelMessageSend(msg.ChannelID, "Error fetching user try again")
				logger.Log.Error(err.Error())
				break
			}
		}

		for _, anime := range animeList {
			if _, exists := animePresentLookup[anime]; !exists {
				id, err := aRepo.Create(anime)
				if err != nil {
					session.ChannelMessageSend(msg.ChannelID, "Error creating anime try again")
					logger.Log.Error(err.Error())
					break
				}
				err = sRepo.Create(models.Subscription{DiscordID: msg.Author.ID, AnimeID: id})
				if err != nil {
					session.ChannelMessageSend(msg.ChannelID, "Error subscribing")
					logger.Log.Error(err.Error())
					break
				}
			} else {
				anime, err := aRepo.GetByTitle(anime)
				if err != nil {
					session.ChannelMessageSend(msg.ChannelID, "Error subscribing")
					logger.Log.Error((err.Error()))
				}
				err = sRepo.Create(models.Subscription{DiscordID: msg.Author.ID, AnimeID: anime.ID})
				if err != nil {
					session.ChannelMessageSend(msg.ChannelID, "Error subscribing")
					logger.Log.Error(err.Error())
					break
				}
			}
		}
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

}
