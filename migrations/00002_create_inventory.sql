-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS inventory(
    inventory_id BIGSERIAL PRIMARY KEY,
    inventory_quantity BIGINT NOT NULL,
    inventory_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    inventory_updated_at TIMESTAMP DEFAULT NULL,
    inventory_deleted_at TIMESTAMP DEFAULT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DROP TABLE IF EXISTS inventory;
-- +goose StatementEnd
