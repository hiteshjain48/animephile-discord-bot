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
	DBUrl     string
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
	DBUrl = os.Getenv("DB_URL")
	if Token == "" {
		return errors.New("no token found")
	}
	if BotPrefix == "" {
		return errors.New("no bot prefix")
	}

	if DBUrl == "" {
		DBUrl = "postgresql://postgres:HesAyEydKMQJdRvDGnVBEZZinGicaRdx@postgres.railway.internal:5432/railway?sslmode=disable"
	}
	// if err != nil {
	// 	return err
	// }
	return nil
}
