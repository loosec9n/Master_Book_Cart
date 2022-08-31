-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS products_user(
    product_user_id BIGSERIAL PRIMARY KEY,
    product_user_name VARCHAR(400) NOT NULL,
    product_user_count BIGINT,
    product_user_descripition VARCHAR(500),
    product_user_author BIGINT REFERENCES product_author(author_id),
    product_user_price NUMERIC,
    product_user_rating NUMERIC,
    product_user_image VARCHAR(500),
    product_user_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    product_user_updated_at  TIMESTAMP NULL,
    product_user_deleted_at TIMESTAMP NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DROP TABLE IF EXISTS product_user CASCADE; 
-- +goose StatementEnd
