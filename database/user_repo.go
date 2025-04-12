package database

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
		"INSERT INTO users (id, name) VALUES ($1 $2)", user.ID, user.Name,
	)
}

func (ur *UserRepository) GetByID(id string) (models.User, error) {
	var user models.User
	row := ur.repo.QueryRow("SELECT id, name FROM users WHERE id = $1", id)

	err := row.Scan(&user.ID, &user.Name)
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
		"SELECT id, name from users",
		func(rows *sql.Rows) error {
			for rows.Next() {
				var user models.User
				if err := rows.Scan(&user.ID, &user.Name); err != nil {
					return err
				}
				users = append(users, user)
			}
			return rows.Err()
		},
	)
	return users, err
}
