-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_order (
    order_id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(user_id),
    address_id BIGINT REFERENCES user_address(user_address_id),
    product_id BIGINT REFERENCES product(product_id),
    inventory_id BIGINT REFERENCES inventory(inventory_id),
    quantity BIGINT NOT NULL,
    product_price NUMERIC NOT NULL,
    total_price NUMERIC NOT NULL,
    order_status VARCHAR(100) DEFAULT 'ordered',
    order_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DROP TABLE IF EXISTS user_order CASCADE;
-- +goose StatementEnd
