-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS wishlist (
    wishlist_id BIGSERIAL PRIMARY KEY,
    product_id BIGINT REFERENCES product(product_id),
    user_id BIGINT REFERENCES users(user_id),
    product_image VARCHAR(100),
    product_price NUMERIC,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DROP TABLE IF EXISTS wishlist CASCADE;
-- +goose StatementEnd
