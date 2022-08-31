-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
      user_id BIGSERIAL PRIMARY KEY,
      is_active BOOLEAN DEFAULT TRUE,
      first_name VARCHAR(200) NOT NULL,
      last_name VARCHAR(200) NOT NULL,
      password VARCHAR(200) NOT NULL,
      email VARCHAR(200) NOT NULL,
      phone_number BIGINT,
      is_admin BOOLEAN DEFAULT FALSE,
      user_address_id BIGINT REFERENCES user_address (user_address_id),
      order_status BIGINT,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP DEFAULT NULL,
      deleted_at TIMESTAMP NULL
   );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DROP TABLE  IF EXISTS users CASCADE;
-- +goose StatementEnd


