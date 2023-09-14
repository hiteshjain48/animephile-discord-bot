package config

import (
	"errors"
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
	Token = os.Getenv("API_KEY")
	BotPrefix = os.Getenv("BOT_PREFIX")
	if Token == ""{
		return errors.New("no token found")
	}
	if BotPrefix == "" {
		return errors.New("no bot prefix")
	}

	return nil
}