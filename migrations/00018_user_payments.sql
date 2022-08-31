-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT  EXISTS user_payment(
    payment_id BIGSERIAL PRIMARY KEY,
    cod_payment BOOLEAN,
    user_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DROP TABLE IF EXISTS user_payments CASCADE;
-- +goose StatementEnd 





