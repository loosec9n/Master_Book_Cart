-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS product(
    product_id BIGSERIAL PRIMARY KEY,
    product_name VARCHAR(400) NOT NULL,
    product_description VARCHAR(500),
    product_category_id BIGINT REFERENCES product_category(category_id),
    product_author_id BIGINT REFERENCES product_author(author_id),
    product_price NUMERIC,
    product_created_at TIMESTAMP DEFAULT NOW(),
    product_updated_at TIMESTAMP,
    product_deleted_at TIMESTAMP
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE product CASCADE;
-- +goose StatementEnd
