-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cart(
    cart_id BIGSERIAL PRIMARY KEY,
    -- session_id BIGINT NOT NULL REFERENCES session_cart(session_id),
    user_id BIGINT NOT NULL REFERENCES users(user_id),
    product_id BIGINT NOT NULL REFERENCES product(product_id),
    product_count BIGINT DEFAULT 10,
    cart_created_at TIMESTAMP DEFAULT NOW(),
    cart_updated_at TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cart;
-- +goose StatementEnd
