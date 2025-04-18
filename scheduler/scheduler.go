package scheduler

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/hiteshjain48/animephile-discord-bot/anime"
	"github.com/hiteshjain48/animephile-discord-bot/database/repositories"
	"github.com/hiteshjain48/animephile-discord-bot/logger"
	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron          *cron.Cron
	animeRepo     *repositories.AnimeRepository
	subRepo       *repositories.SubscriptionRepository
	discordClient *discordgo.Session
}

func NewScheduler(
	animeRepo *repositories.AnimeRepository,
	subRepo *repositories.SubscriptionRepository,
	discord *discordgo.Session,
) *Scheduler {
	c := cron.New(cron.WithSeconds())

	return &Scheduler{
		cron:          c,
		animeRepo:     animeRepo,
		subRepo:       subRepo,
		discordClient: discord,
	}
}

func (s *Scheduler) Start() {
	_, err := s.cron.AddFunc("0 0 * * * *", s.checkNewEpisodes)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Error starting scheduler: %v", err.Error()))
		return
	}

	s.cron.Start()
	logger.Log.Info("Scheduler Started")
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
}

func (s *Scheduler) checkNewEpisodes() {
	logger.Log.Info("Checking for new episodes...")

	schedule, err := anime.GetSchedule()
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Error while getting schedule from api: %v", err.Error()))
	}

	subscriptions, err := s.subRepo.List()
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Error while getting subscriptions from db: %v", err.Error()))
	}

	animeMap := make(map[int][]string)
	for _, sub := range subscriptions {
		animeMap[sub.AnimeID] = append(animeMap[sub.AnimeID], sub.DiscordID)
	}

	for animeID, subscribers := range animeMap {
		anime, err := s.animeRepo.GetByID(animeID)
		if err != nil {
			logger.Log.Error(fmt.Sprintf("Error: %v", err.Error()))
			continue
		}
		for _, airedAnime := range schedule {
			if anime.Title == airedAnime.Media.Title.Romaji {
				for _, discordID := range subscribers {
					s.notifyUser(discordID, airedAnime)
				}
			}
		}
	}
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Error fetching anime from db: %v", err.Error()))
	}
}

func (s *Scheduler) notifyUser(discordID string, airedAnime anime.AiringSchedule) {
	channel, err := s.discordClient.UserChannelCreate(discordID)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Error creating user channel: %v", err.Error()))
		return
	}

	msg := fmt.Sprintf("📺 %s - Ep %d at %s\n",
		airedAnime.Media.Title.Romaji,
		airedAnime.Episode,
		time.Unix(airedAnime.AiringAt, 0).Format("15:04 MST"))

	_, err = s.discordClient.ChannelMessageSend(channel.ID, msg)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("Error sending notification to user %s: %v", discordID, err))
	}
}
