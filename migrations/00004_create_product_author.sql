-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS product_author(
    author_id BIGSERIAL PRIMARY KEY,
    author_name VARCHAR(400) NOT NULL,
    author_created_at TIMESTAMP DEFAULT NOW(),
    author_updated_at TIMESTAMP DEFAULT NOW(),
    author_deleted_at TIMESTAMP DEFAULT NOW()
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS product_author CASCADE;
-- +goose StatementEnd
