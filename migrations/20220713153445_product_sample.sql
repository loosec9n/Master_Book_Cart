-- +goose Up
-- +goose StatementBegin
INSERT INTO product (
		product_name, 
		product_description, 
		product_price,
		product_author_id,
		product_category_id)
Values
    ('Book1','First Book',110.0,201,101),
    ('Book2','Second Book',120.1,202,102),
    ('Book3','Third Book',130.2,203,103),
    ('Book4','Fourth Book',140.3,204,104),
    ('Book5','Fifth Book',150.4,205,105),
    ('Book6','Sixth Book',160.5,201,106),
    ('Book7','Seventh Book',170.6,202,107),
    ('Book8','Eighth Book',180.7,203,108),
    ('Book9','Nineth Book',190.8,204,109),
    ('Book10','Thenth Book',200.9,205,101),
    ('Book11','Eleventh Book',210.0,201,102),
    ('Book12','Twelth Book',220.1,202,103),
    ('Book13','Thirteenth Book',230.2,203,104),
    ('Book14','Fourteenth Book',240.3,204,105),
    ('Book15','Fifteenth Book',250.4,205,106),
    ('Book16','Sixteenth Book',260.5,201,107),
    ('Book17','Seventeenth Book',270.6,202,108),
    ('Book18','Eighteenth Book',280.7,203,109),
    ('Book19','Ninteenth Book',290.9,204,101),
    ('Book20','Twenteenth Book',300.0,205,102);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE product CASCADE;
-- +goose StatementEnd
