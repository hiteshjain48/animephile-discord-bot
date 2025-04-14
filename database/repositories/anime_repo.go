package repositories

import(
	"database/sql"
	"errors"

	"github.com/hiteshjain48/animephile-discord-bot/database/models"
)

type AnimeRepository struct {
	repo *Repository
}

func NewAnimeRepository(db *sql.DB) *AnimeRepository {
	return &AnimeRepository{
		repo: NewRepository(db),
	}
}

func (ar *AnimeRepository) Create(anime models.Anime) error {
	return ar.repo.Execute(
		"INSERT INTO anime (id, title) VALUES ($1, $2)", anime.ID, anime.Title,
	)
}

func (ar *AnimeRepository) GetByID(id string) (models.Anime, error) {
	var anime models.Anime
	row := ar.repo.QueryRow("SELECT id, title FROM anime WHERE id = $1", id)

	err := row.Scan(&anime.ID, &anime.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			return anime, errors.New("anime not found")
		}
		return anime, err
	}

	return anime, nil
}

func (ar *AnimeRepository) GetByTitle(title string) (models.Anime, error) {
	var anime models.Anime
	row := ar.repo.QueryRow("SELECT id, title FROM anime WHERE title = $1", title)

	err := row.Scan(&anime.ID, &anime.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			return anime, errors.New("anime not found")
		}
		return anime, err
	}

	return anime, nil
}

func (ar *AnimeRepository) List() ([]models.Anime, error) {
	var animes []models.Anime

	err := ar.repo.Query(
		"SELECT id, title from anime",
		func(rows *sql.Rows) error {
			for rows.Next() {
				var anime models.Anime
				if err := rows.Scan(&anime.ID, &anime.Title); err != nil {
					return err
				}
				animes = append(animes, anime)
			}
			return rows.Err()
		},
	)
	return animes, err
}