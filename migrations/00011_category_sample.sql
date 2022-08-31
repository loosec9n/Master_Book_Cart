-- +goose Up
-- +goose StatementBegin
INSERT INTO 
    product_category (
		category_id, 
        category_name, 
        category_description)
VALUES
    (101,'Drama','book of dramas'),
    (102,'Horror','book of horror'),
    (103,'Classics','book of classics novels'),
    (104,'Tragedy','collection of tragedy novels'),
    (105,'Sci-Fi','collection of unbelievable imaginations'),
    (106,'Fantasy','collecton of fantasy books'),
    (107,'Comics','Book of comics'),
    (108,'Romance','Collection of books on romance stories'),
    (109,'Biography','collection of motivational books'),
    (110,'Crime','collection of thrilling findings');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- TRUNCATE TABLE product_category CASCADE;
-- +goose StatementEnd
