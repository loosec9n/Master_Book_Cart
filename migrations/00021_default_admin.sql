-- +goose Up
-- +goose StatementBegin

INSERT INTO users(first_name, last_name, password, email, phone_number, is_admin)
VALUES ('Justin', 'John', '5f4dcc3b5aa765d61d8327deb882cf99', 'justin@admin.com', 9539598855, true );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DELETE FROM users WHERE email = 'justin@admin.com';
-- +goose StatementEnd