package repositories

import (
	"database/sql"
	"errors"

	"github.com/hiteshjain48/animephile-discord-bot/database/models"
)

type UserRepository struct {
	repo *Repository
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		repo: NewRepository(db),
	}
}

func (ur *UserRepository) Create(user models.User) error {
	return ur.repo.Execute(
		"INSERT INTO users (discord_id, user_name, joined_at) VALUES ($1, $2, $3) ON CONFLICT (discord_id) DO NOTHING", user.DiscordID, user.UserName, user.JoinedAt,
	)
}

func (ur *UserRepository) GetByID(id string) (models.User, error) {
	var user models.User
	row := ur.repo.QueryRow("SELECT discord_id, user_name FROM users WHERE discord_id = $1", id)

	err := row.Scan(&user.DiscordID, &user.UserName)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user not found")
		}
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) List() ([]models.User, error) {
	var users []models.User

	err := ur.repo.Query(
		"SELECT discord_id, user_name from users",
		func(rows *sql.Rows) error {
			for rows.Next() {
				var user models.User
				if err := rows.Scan(&user.DiscordID, &user.UserName); err != nil {
					return err
				}
				users = append(users, user)
			}
			return rows.Err()
		},
	)
	return users, err
}


