package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"	
)

var (
	Token string
	BotPrefix string
)

func ReadConfig() (error) {
	fmt.Println("Reading .env file")
	godotenv.Load(".env")
	Token = os.Getenv("TOKEN")
	BotPrefix = os.Getenv("BOT_PREFIX")
	if Token == ""{
		return error.New("no api found")
	}
	if BotPrefix == "" {
		return error.New("no bot prefix")
	}

	return nil
}