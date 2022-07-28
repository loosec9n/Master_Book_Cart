-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_order_status (
    order_id BIGINT PRIMARY KEY,
    order_price NUMERIC(5,2) NOT NULL,
    order_discount INTEGER,
    payment_method BIGINT REFERENCES user_payment(payment_id),
    order_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_order_status;
-- +goose StatementEnd




