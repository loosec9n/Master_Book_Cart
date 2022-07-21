-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS discount(
    discount_id BIGSERIAL PRIMARY KEY,
    discount_name VARCHAR(400) NOT NULL,
    discount_description VARCHAR(500),
    discount_percentage NUMERIC(2,2),
    discount_status BOOLEAN NOT NULL,
    discount_created_at TIMESTAMP DEFAULT NOW(),
    discount_updates_at TIMESTAMP,
    discount_deleted_at TIMESTAMP
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS discount;
-- +goose StatementEnd
