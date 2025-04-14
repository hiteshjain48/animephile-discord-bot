package models

import(
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID           uuid.UUID       
	DiscordID    string    
	AnimeID      int       
	SubscribedAt time.Time 
}