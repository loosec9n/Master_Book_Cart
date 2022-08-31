-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS product(
    product_id BIGSERIAL PRIMARY KEY,
    is_active BOOLEAN DEFAULT TRUE,
    product_name VARCHAR(400) NOT NULL,
    product_description VARCHAR(500),
    product_category_id BIGINT REFERENCES product_category(category_id),
    product_author_id BIGINT REFERENCES product_author(author_id),
    product_inventory_id BIGINT REFERENCES inventory(inventory_id),
    product_price NUMERIC,
    product_rating NUMERIC,
    product_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    product_updated_at TIMESTAMP DEFAULT NULL,
    product_deleted_at TIMESTAMP DEFAULT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DROP TABLE product CASCADE;
-- +goose StatementEnd
