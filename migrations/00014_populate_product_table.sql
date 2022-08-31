-- +goose Up
-- +goose StatementBegin
INSERT INTO product (
		product_name, 
		product_description, 
		product_price,
		product_author_id,
		product_category_id,
        product_inventory_id,
        product_rating)
Values
    ('Book1','First Book',110.0,201,101,301,3.8),
    ('Book2','Second Book',120.1,202,102,302,3.9),
    ('Book3','Third Book',130.2,203,103,303,3.7),
    ('Book4','Fourth Book',140.3,204,104,304,3.9),
    ('Book5','Fifth Book',150.4,205,105,305,4.0),
    ('Book6','Sixth Book',160.5,201,106,306,4.1),
    ('Book7','Seventh Book',170.6,202,107,307,4.2),
    ('Book8','Eighth Book',180.7,203,108,308,4.3),
    ('Book9','Nineth Book',190.8,204,109,309,4.4),
    ('Book10','Thenth Book',200.9,205,101,310,4.5),
    ('Book11','Eleventh Book',210.0,201,102,311,4.6),
    ('Book12','Twelth Book',220.1,202,103,312,4.7),
    ('Book13','Thirteenth Book',230.2,203,104,313,4.8),
    ('Book14','Fourteenth Book',240.3,204,105,314,4.9),
    ('Book15','Fifteenth Book',250.4,205,106,315,4.5),
    ('Book16','Sixteenth Book',260.5,201,107,316,4.4),
    ('Book17','Seventeenth Book',270.6,202,108,317,4.6),
    ('Book18','Eighteenth Book',280.7,203,109,318,4.0),
    ('Book19','Ninteenth Book',290.9,204,101,319,4.1),
    ('Book20','Twenteenth Book',300.0,205,102,320,4.4);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- TRUNCATE TABLE product CASCADE;
-- +goose StatementEnd
