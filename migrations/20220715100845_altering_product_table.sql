-- +goose Up
-- +goose StatementBegin
ALTER TABLE product
ADD is_active BOOLEAN DEFAULT TRUE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE product
DROP COLUMN is_active;
-- +goose StatementEnd
