-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_order (
    order_id BIGINT PRIMARY KEY,
    user_id BIGINT REFERENCES users(user_id),
    address_id BIGINT REFERENCES user_address(user_address_id),
    product_id BIGINT REFERENCES product(product_id),
    inventory_id BIGINT REFERENCES inventory(inventory_id),
    quantity BIGINT NOT NULL,
    cart_id BIGINT REFERENCES cart(cart_id) ON DELETE CASCADE,
    total_price NUMERIC,
    order_status VARCHAR(100) DEFAULT 'ordered',
    order_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_order CASCADE;
-- +goose StatementEnd
