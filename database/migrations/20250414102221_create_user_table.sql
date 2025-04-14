-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    discord_id VARCHAR(20) PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    joined_at TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
