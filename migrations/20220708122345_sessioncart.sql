-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS session_cart(
    session_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(user_id),
    session_created_at TIMESTAMP DEFAULT NOW(),
    session_updated_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS session_cart;
-- +goose StatementEnd
