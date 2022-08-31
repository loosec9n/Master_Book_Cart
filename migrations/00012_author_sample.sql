-- +goose Up
-- +goose StatementBegin
INSERT INTO 
    product_author (
		author_id, 
        author_name)
VALUES
    (201,'Bimal Jalal'),
    (202,'Ruskin Bond'),
    (203,'Vinit Karnik'),
    (204,'Preeti Shenoy'),
    (205,'Smriti Irani');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- TRUNCATE TABLE product_author CASCADE;
-- +goose StatementEnd
