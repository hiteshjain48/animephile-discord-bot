-- +goose Up
-- +goose StatementBegin
CREATE TABLE subscriptions (
    id SERIAL PRIMARY KEY,
    discord_id VARCHAR(20) REFERENCES users(discord_id) ON DELETE CASCADE,
    anime_id INTEGER REFERENCES anime(id) ON DELETE CASCADE,
    UNIQUE(discord_id, anime_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS subscriptions;
-- +goose StatementEnd
