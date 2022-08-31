-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS ordered(
    order_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    total_amount NUMERIC NOT NULL,
    status BOOLEAN DEFAULT FALSE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DROP TABLE IF EXISTS ordered CASCADE;
-- +goose StatementEnd
