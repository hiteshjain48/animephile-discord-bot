package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/hiteshjain48/animephile-discord-bot/logger"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
	DBUrl	 string
}

func Connect(config DBConfig) (*sql.DB, error) {
	// connStr := fmt.Sprintf(
	// 	"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
	// 	config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	// )
	connStr := config.DBUrl
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func RunMigrations(db *sql.DB, migrationPath string) error {
	goose.SetBaseFS(nil)
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db, migrationPath); err != nil {
		return err
	}

	logger.Log.Info("Migrations successfully applied")
	return nil
}
