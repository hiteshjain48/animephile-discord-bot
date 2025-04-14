package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Token     string
	BotPrefix string
	DBHost    string
	DBPort    int
	DBUser    string
	DBPass    string
	DBName    string
	DBSSLMode string
)

func ReadConfig() error {
	fmt.Println("Reading .env file")
	godotenv.Load(".env")
	var err error
	Token = os.Getenv("API_KEY")
	BotPrefix = os.Getenv("BOT_PREFIX")
	DBHost = os.Getenv("DB_HOST")
	DBPort, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		DBPort = 5432
	}
	DBName = os.Getenv("DB_NAME")
	DBUser = os.Getenv("DB_USER")
	DBPass = os.Getenv("DB_PASS")
	DBSSLMode = os.Getenv("DB_SSL_MODE")
	if Token == "" {
		return errors.New("no token found")
	}
	if BotPrefix == "" {
		return errors.New("no bot prefix")
	}
	// if err != nil {
	// 	return err
	// }
	return nil
}
