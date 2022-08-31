-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_address(
user_address_id BIGSERIAL PRIMARY KEY,
address_id INTEGER NOT NULL,
house_name VARCHAR(400) NOT NULL,
street_name VARCHAR(400) NOT NULL,
land_mark VARCHAR(400),
city VARCHAR(200) NOT NULL,
add_state VARCHAR(200) NOT NULL,
pincode INTEGER NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DROP TABLE IF EXISTS user_address;
-- +goose StatementEnd
