-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS product_category(
    category_id BIGSERIAL PRIMARY KEY,
    category_name VARCHAR(400) NOT NULL,
    category_description VARCHAR(500),
    category_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    category_updated_at TIMESTAMP DEFAULT NOW(),
    category_deleted_at TIMESTAMP DEFAULT NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DROP TABLE IF EXISTS product_category CASCADE;
-- +goose StatementEnd
