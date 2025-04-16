package repositories


import(
	"database/sql"
	"errors"

	"github.com/hiteshjain48/animephile-discord-bot/database/models"
)

type SubscriptionRepository struct {
	repo *Repository
}

func NewSubscriptionRepository(db *sql.DB) *SubscriptionRepository {
	return &SubscriptionRepository{
		repo: NewRepository(db),
	}
}

func (sr *SubscriptionRepository) Create(subscription models.Subscription) error {
	return sr.repo.Execute(
		"INSERT INTO subscriptions (discord_id, anime_id) VALUES ($1, $2)", subscription.DiscordID, subscription.AnimeID,
	)
}

func (sr *SubscriptionRepository) GetByID(id string) (models.Subscription, error) {
	var subscription models.Subscription
	row := sr.repo.QueryRow("SELECT discord_id, anime_id FROM anime WHERE id = $1", id)

	err := row.Scan(&subscription.DiscordID, &subscription.AnimeID)
	if err != nil {
		if err == sql.ErrNoRows {
			return subscription, errors.New("subscription not found")
		}
		return subscription, err
	}

	return subscription, nil
}



func (sr *SubscriptionRepository) List() ([]models.Subscription, error) {
	var subscriptions []models.Subscription

	err := sr.repo.Query(
		"SELECT discord_id, anime_id from subscription",
		func(rows *sql.Rows) error {
			for rows.Next() {
				var subscription models.Subscription
				if err := rows.Scan(&subscription.DiscordID, &subscription.AnimeID); err != nil {
					return err
				}
				subscriptions = append(subscriptions, subscription)
			}
			return rows.Err()
		},
	)
	return subscriptions, err
}