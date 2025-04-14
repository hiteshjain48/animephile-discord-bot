package models

import (
	"time"
)

type User struct {
	DiscordID 	string
	UserName 	string
	JoinedAt 	time.Time
}

